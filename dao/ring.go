/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package dao

import (
	"github.com/expanse-org/relay-lib/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	"errors"
	"math/big"
	"time"
	"github.com/expanse-org/relay-lib/log"
)

type FilledOrder struct {
	ID               int    `gorm:"column:id;primary_key;"`
	RingHash         string `gorm:"column:ringhash;type:varchar(82)"`
	OrderHash        string `gorm:"column:orderhash;type:varchar(82)"`
	FeeSelection     uint8  `gorm:"column:fee_selection" json:"feeSelection"`
	RateAmountS      string `gorm:"column:rate_amount_s;type:text" json:"rateAmountS"`
	AvailableAmountS string `gorm:"column:available_amount_s;type:text"json:"availableAmountS"`
	AvailableAmountB string `gorm:"column:available_amount_b;type:text"`
	FillAmountS      string `gorm:"column:fill_amount_s;type:text" json:"fillAmountS"`
	FillAmountB      string `gorm:"column:fill_amount_b;type:text" json:"fillAmountB"`
	LrcReward        string `gorm:"column:lrc_reward;type:text" json:"lrcReward"`
	LrcFee           string `gorm:"column:lrc_fee;type:text" json:"lrcFee"`
	FeeS             string `gorm:"column:fee_s;type:text" json:"feeS"`
	LegalFee         string `gorm:"column:legal_fee;type:text" json:"legalFee"`
	SPrice           string `gorm:"column:s_price;type:text" json:"sPrice"`
	BPrice           string `gorm:"column:b_price;type:text" json:"sPrice"`
}

func getRatString(v *big.Rat) string {
	if nil == v {
		return ""
	} else {
		return v.String()
	}
}

func (daoFilledOrder *FilledOrder) ConvertDown(filledOrder *types.FilledOrder, ringhash common.Hash) error {
	daoFilledOrder.RingHash = ringhash.Hex()
	daoFilledOrder.OrderHash = filledOrder.OrderState.RawOrder.Hash.Hex()
	daoFilledOrder.FeeSelection = filledOrder.FeeSelection
	daoFilledOrder.RateAmountS = getRatString(filledOrder.RateAmountS)
	daoFilledOrder.AvailableAmountS = getRatString(filledOrder.AvailableAmountS)
	daoFilledOrder.AvailableAmountB = getRatString(filledOrder.AvailableAmountB)
	daoFilledOrder.FillAmountS = getRatString(filledOrder.FillAmountS)
	daoFilledOrder.FillAmountB = getRatString(filledOrder.FillAmountB)
	daoFilledOrder.LrcReward = getRatString(filledOrder.LrcReward)
	daoFilledOrder.LrcFee = getRatString(filledOrder.LrcFee)
	daoFilledOrder.FeeS = getRatString(filledOrder.FeeS)
	daoFilledOrder.LegalFee = getRatString(filledOrder.LegalFee)
	daoFilledOrder.SPrice = getRatString(filledOrder.SPrice)
	daoFilledOrder.BPrice = getRatString(filledOrder.BPrice)
	return nil
}

func (daoFilledOrder *FilledOrder) ConvertUp(filledOrder *types.FilledOrder, rds RdsService) error {
	//if nil != rds {
	//	daoOrderState, err := rds.GetOrderByHash(common.HexToHash(daoFilledOrder.OrderHash))
	//	if nil != err {
	//		return err
	//	}
	//	orderState := &types.OrderState{}
	//	daoOrderState.ConvertUp(orderState)
	//	filledOrder.OrderState = *orderState
	//}
	filledOrder.FeeSelection = daoFilledOrder.FeeSelection
	filledOrder.RateAmountS = new(big.Rat)
	filledOrder.RateAmountS.SetString(daoFilledOrder.RateAmountS)
	filledOrder.AvailableAmountS = new(big.Rat)
	filledOrder.AvailableAmountB = new(big.Rat)
	filledOrder.AvailableAmountS.SetString(daoFilledOrder.AvailableAmountS)
	filledOrder.AvailableAmountB.SetString(daoFilledOrder.AvailableAmountB)
	filledOrder.FillAmountS = new(big.Rat)
	filledOrder.FillAmountB = new(big.Rat)
	filledOrder.FillAmountS.SetString(daoFilledOrder.FillAmountS)
	filledOrder.FillAmountB.SetString(daoFilledOrder.FillAmountB)
	filledOrder.LrcReward = new(big.Rat)
	filledOrder.LrcFee = new(big.Rat)
	filledOrder.LrcReward.SetString(daoFilledOrder.LrcReward)
	filledOrder.LrcFee.SetString(daoFilledOrder.LrcFee)
	filledOrder.FeeS = new(big.Rat)
	filledOrder.FeeS.SetString(daoFilledOrder.FeeS)
	filledOrder.LegalFee = new(big.Rat)
	filledOrder.LegalFee.SetString(daoFilledOrder.LegalFee)
	filledOrder.SPrice = new(big.Rat)
	filledOrder.SPrice.SetString(daoFilledOrder.SPrice)
	filledOrder.BPrice = new(big.Rat)
	filledOrder.BPrice.SetString(daoFilledOrder.BPrice)
	return nil
}

func (s *RdsServiceImpl) GetFilledOrderByRinghash(ringhash common.Hash) ([]*FilledOrder, error) {
	var (
		filledOrders []*FilledOrder
		err          error
	)

	err = s.Db.Where("ringhash = ?", ringhash.Hex()).
		Find(&filledOrders).
		Error

	return filledOrders, err
}

type RingSubmitInfo struct {
	ID               int       `gorm:"column:id;primary_key;"`
	RingHash         string    `gorm:"column:ringhash;type:varchar(82)"`
	UniqueId         string    `gorm:"column:unique_id;type:varchar(82)"`
	ProtocolAddress  string    `gorm:"column:protocol_address;type:varchar(42)"`
	OrdersCount      int64     `gorm:"column:order_count;type:bigint"`
	ProtocolData     string    `gorm:"column:protocol_data;type:text"`
	ProtocolGas      string    `gorm:"column:protocol_gas;type:varchar(50)"`
	ProtocolGasPrice string    `gorm:"column:protocol_gas_price;type:varchar(50)"`
	ProtocolUsedGas  string    `gorm:"column:protocol_used_gas;type:varchar(50)"`
	ProtocolTxHash   string    `gorm:"column:protocol_tx_hash;type:varchar(82)"`
	TxNonce          uint64    `gorm:"column:tx_nonce;type:bigint"`
	Status           int       `gorm:"column:status;type:int"`
	RingIndex        string    `gorm:"column:ring_index;type:varchar(50)"`
	BlockNumber      string    `gorm:"column:block_number;type:varchar(50)"`
	Miner            string    `gorm:"column:miner;type:varchar(42)"`
	Err              string    `gorm:"column:err;type:text"`
	CreateTime       time.Time `gorm:"column:create_time;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
}

func getBigIntString(v *big.Int) string {
	if nil == v {
		return ""
	} else {
		return v.String()
	}
}

func (info *RingSubmitInfo) ConvertDown(typesInfo *types.RingSubmitInfo, err error) error {
	info.RingHash = typesInfo.Ringhash.Hex()
	info.UniqueId = typesInfo.RawRing.GenerateUniqueId().Hex()
	info.ProtocolAddress = typesInfo.ProtocolAddress.Hex()
	info.OrdersCount = typesInfo.OrdersCount.Int64()
	info.ProtocolData = common.ToHex(typesInfo.ProtocolData)
	info.ProtocolGas = getBigIntString(typesInfo.ProtocolGas)
	info.ProtocolUsedGas = getBigIntString(typesInfo.ProtocolUsedGas)
	info.ProtocolGasPrice = getBigIntString(typesInfo.ProtocolGasPrice)
	info.Miner = typesInfo.Miner.Hex()
	info.ProtocolTxHash = typesInfo.SubmitTxHash.Hex()
	if nil != err {
		info.Err = err.Error()
	}
	return nil
}

func (info *RingSubmitInfo) ConvertUp(typesInfo *types.RingSubmitInfo) error {
	typesInfo.Ringhash = common.HexToHash(info.RingHash)
	typesInfo.ProtocolAddress = common.HexToAddress(info.ProtocolAddress)
	typesInfo.OrdersCount = big.NewInt(info.OrdersCount)
	typesInfo.ProtocolData = common.FromHex(info.ProtocolData)
	typesInfo.ProtocolGas = new(big.Int)
	typesInfo.ProtocolGas.SetString(info.ProtocolGas, 0)
	typesInfo.ProtocolUsedGas = new(big.Int)
	typesInfo.ProtocolUsedGas.SetString(info.ProtocolUsedGas, 0)
	typesInfo.ProtocolGasPrice = new(big.Int)
	typesInfo.ProtocolGasPrice.SetString(info.ProtocolGasPrice, 0)
	typesInfo.SubmitTxHash = common.HexToHash(info.ProtocolTxHash)
	typesInfo.Miner = common.HexToAddress(info.Miner)
	return nil
}

//func (s *RdsServiceImpl) UpdateRingSubmitInfoRegistryTxHash(ringhashs []common.Hash, txHash string) error {
//	hashes := []string{}
//	for _, h := range ringhashs {
//		hashes = append(hashes, h.Hex())
//	}
//	dbForUpdate := s.db.Model(&RingSubmitInfo{}).Where("ringhash in (?)", hashes)
//	return dbForUpdate.Update("registry_tx_hash", txHash).Error
//}

//func (s *RdsServiceImpl) UpdateRingSubmitInfoFailed(ringhashs []common.Hash, err string) error {
//	hashes := []string{}
//	for _, h := range ringhashs {
//		hashes = append(hashes, h.Hex())
//	}
//	dbForUpdate := s.db.Model(&RingSubmitInfo{}).Where("ringhash in (?) ", hashes)
//	return dbForUpdate.Update("err", err).Error
//}

func (s *RdsServiceImpl) UpdateRingSubmitInfoResult(submitResult *types.RingSubmitResultEvent) error {
	items := map[string]interface{}{
		"status":            uint8(submitResult.Status),
		"ring_index":        getBigIntString(submitResult.RingIndex),
		"block_number":      getBigIntString(submitResult.BlockNumber),
		"protocol_used_gas": getBigIntString(submitResult.UsedGas),
		"ringhash":          submitResult.RingHash.Hex(),
		"protocol_tx_hash":  submitResult.TxHash.Hex(),
		"err":               submitResult.Err,
	}
	if submitResult.TxNonce > 0 {
		items["tx_nonce"] = submitResult.TxNonce
	}
	//if "" != submitResult.Err {
	//	items["err"] = submitResult.Err
	//}

	var dbForUpdate *gorm.DB
	if submitResult.RecordId > 0 {
		dbForUpdate = s.Db.Model(&RingSubmitInfo{}).Where("id = ?", submitResult.RecordId)
	} else {
		dbForUpdate = s.Db.Model(&RingSubmitInfo{}).Where("ringhash = ? and protocol_tx_hash = ? ", submitResult.RingHash.Hex(), submitResult.TxHash.Hex())
	}
	//todo:test it

	return dbForUpdate.Update(items).Error
}

//func (s *RdsServiceImpl) UpdateRingSubmitInfoProtocolTxHash(ringhash common.Hash, txHash string) error {
//	dbForUpdate := s.db.Model(&RingSubmitInfo{}).Where("ringhash = ?", ringhash.Hex())
//	return dbForUpdate.Update("protocol_tx_hash", txHash).Error
//}

func (s *RdsServiceImpl) GetRingForSubmitByHash(ringhash common.Hash) (ringForSubmit RingSubmitInfo, err error) {
	err = s.Db.Where("ringhash = ? ", ringhash.Hex()).First(&ringForSubmit).Error
	return
}

func (s *RdsServiceImpl) HasReSubmited(createTime int64, miner string, txNonce uint64) (bool, error) {
	t := big.NewInt(createTime)
	count := 0
	err := s.Db.Model(&RingSubmitInfo{}).Where("UNIX_TIMESTAMP(create_time) > ? and miner = ? and tx_nonce = ?", t.String(), miner, txNonce).Count(&count).Error
	return (count>0),err
}
func (s *RdsServiceImpl) GetPendingTx(createTime int64) (ringForSubmits []RingSubmitInfo, err error) {
	ringForSubmits = []RingSubmitInfo{}
	t := big.NewInt(createTime)
	if err := s.Db.Raw("select infos.* from lpr_ring_submit_infos as infos " +
	" join " +
	" (select miner, max(tx_nonce) blockedNonce from lpr_ring_submit_infos  where status in (2,3) group by miner) as blockedNonces on blockedNonces.miner=infos.miner and status = 1 " +
	"	and infos.tx_nonce > blockedNonces.blockedNonce and UNIX_TIMESTAMP(create_time) < " + t.String() + " order by infos.tx_nonce").Scan(&ringForSubmits).Error;nil != err {
		log.Errorf("err:%s", err.Error())
	}
			//minerBlockedNonces := []map[string]interface{}{}
	//if err = s.Db.Raw("select " +
	//	"miner, " +
	//	"max(tx_nonce) blockedNonce " +
	//	" from lpr_ring_submit_infos " +
	//	" where status = 2 " +
	//	" group by miner ").Scan(&minerBlockedNonces).Error; nil == err {
	//	for _, minerBlockNonce := range minerBlockedNonces {
	//		miner := minerBlockNonce["miner"]
	//		nonce := minerBlockNonce["blockedNonce"]
	//		var list []RingSubmitInfo
	//		if err1 := s.Db.Model(&RingSubmitInfo{}).Where(" create_time > ? and status = ? and miner = ? and tx_nonce > ? ", createTime, 0, miner, nonce).Scan(&list).Error; nil == err1 {
	//			if len(list) > 0 {
	//				for _, info := range list {
	//					ringForSubmits = append(ringForSubmits, info)
	//				}
	//			} else {
	//				log.Debugf("can't get pendingtx of owner:%s, nonce:%d in submitringinfo", miner, nonce)
	//			}
	//		} else {
	//			log.Errorf("error:%s", err1.Error())
	//		}
	//	}
	//} else {
	//	log.Errorf("err:%s", err.Error())
	//}
	return
}

func (s *RdsServiceImpl) GetRingHashesByTxHash(txHash common.Hash) ([]*RingSubmitInfo, error) {
	var (
		err   error
		infos []*RingSubmitInfo
	)

	err = s.Db.Where("protocol_tx_hash = ? ", txHash.Hex()).
		Find(&infos).
		Error

	return infos, err
}

func (s *RdsServiceImpl) UpdateRingSubmitInfoSubmitUsedGas(txHash string, usedGas *big.Int) error {
	dbForUpdate := s.Db.Model(&RingSubmitInfo{}).Where("protocol_tx_hash = ?", txHash)
	return dbForUpdate.Update("protocol_used_gas", getBigIntString(usedGas)).Error
}

func (s *RdsServiceImpl) UpdateRingSubmitInfoErrById(id int, err error) error {
	if nil == err {
		err = errors.New("")
	}
	dbForUpdate := s.Db.Model(&RingSubmitInfo{}).Where("id = ?", id)
	return dbForUpdate.Update("err", err.Error()).Error
}

func (s *RdsServiceImpl) GetSubmitterNonce(submitter string) (uint64,error) {
	nonce := []uint64{}
	println(submitter)
	err := s.Db.Model(&RingSubmitInfo{}).Where("miner=?", submitter).Pluck(" max(tx_nonce) ", &nonce).Error
	if len(nonce) > 0 {
		return nonce[0]+1,err
	} else {
		return 0, err
	}
}