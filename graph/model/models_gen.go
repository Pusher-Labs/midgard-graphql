// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Asset struct {
	Asset       string `json:"asset"`
	DateCreated int    `json:"dateCreated"`
	PriceRune   string `json:"priceRune"`
}

type AssetAmount struct {
	Asset  string `json:"asset"`
	Amount string `json:"amount"`
}

type BlockReward struct {
	BlockReward string `json:"blockReward"`
	BondReward  string `json:"bondReward"`
	StakeReward string `json:"stakeReward"`
}

type BondMetrics struct {
	TotalActiveBond    string `json:"totalActiveBond"`
	AverageActiveBond  string `json:"averageActiveBond"`
	MedianActiveBond   string `json:"medianActiveBond"`
	MinimumActiveBond  string `json:"minimumActiveBond"`
	MaximumActiveBond  string `json:"maximumActiveBond"`
	TotalStandbyBond   string `json:"totalStandbyBond"`
	AverageStandbyBond string `json:"averageStandbyBond"`
	MedianStandbyBond  string `json:"medianStandbyBond"`
	MinimumStandbyBond string `json:"minimumStandbyBond"`
	MaximumStandbyBond string `json:"maximumStandbyBond"`
}

type BoolValues struct {
	StrictBondStakeRatio bool `json:"StrictBondStakeRatio"`
}

type Constants struct {
	Int64Values  *Int64Values  `json:"int_64_values"`
	BoolValues   *BoolValues   `json:"bool_values"`
	StringValues *StringValues `json:"string_values"`
}

type Health struct {
	Database      bool `json:"database"`
	ScannerHeight int  `json:"scannerHeight"`
	CatchingUp    bool `json:"catching_up"`
}

type Int64Values struct {
	BadValidatorRate                int `json:"BadValidatorRate"`
	BlocksPerYear                   int `json:"BlocksPerYear"`
	DesireValidatorSet              int `json:"DesireValidatorSet"`
	DoubleSignMaxAge                int `json:"DoubleSignMaxAge"`
	EmissionCurve                   int `json:"EmissionCurve"`
	FailKeySignSlashPoints          int `json:"FailKeySignSlashPoints"`
	FailKeygenSlashPoints           int `json:"FailKeygenSlashPoints"`
	FundMigrationInterval           int `json:"FundMigrationInterval"`
	JailTimeKeygen                  int `json:"JailTimeKeygen"`
	JailTimeKeysign                 int `json:"JailTimeKeysign"`
	LackOfObservationPenalty        int `json:"LackOfObservationPenalty"`
	MinimumBondInRune               int `json:"MinimumBondInRune"`
	MinimumNodesForBft              int `json:"MinimumNodesForBFT"`
	MinimumNodesForYggdrasil        int `json:"MinimumNodesForYggdrasil"`
	NewPoolCycle                    int `json:"NewPoolCycle"`
	ObserveSlashPoints              int `json:"ObserveSlashPoints"`
	OldValidatorRate                int `json:"OldValidatorRate"`
	RotatePerBlockHeight            int `json:"RotatePerBlockHeight"`
	RotateRetryBlocks               int `json:"RotateRetryBlocks"`
	SigningTransactionPeriod        int `json:"SigningTransactionPeriod"`
	StakeLockUpBlocks               int `json:"StakeLockUpBlocks"`
	TransactionFee                  int `json:"TransactionFee"`
	ValidatorRotateInNumBeforeFull  int `json:"ValidatorRotateInNumBeforeFull"`
	ValidatorRotateNumAfterFull     int `json:"ValidatorRotateNumAfterFull"`
	ValidatorRotateOutNumBeforeFull int `json:"ValidatorRotateOutNumBeforeFull"`
	WhiteListGasAsset               int `json:"WhiteListGasAsset"`
	YggFundLimit                    int `json:"YggFundLimit"`
}

type LastBlock struct {
	Chain          string `json:"chain"`
	Lastobservedin string `json:"lastobservedin"`
	Lastsignedout  string `json:"lastsignedout"`
	Thorchain      string `json:"thorchain"`
}

type NetworkData struct {
	BondMetrics             *BondMetrics `json:"bondMetrics"`
	ActiveBonds             []*string    `json:"activeBonds"`
	StandbyBonds            []string     `json:"standbyBonds"`
	TotalStaked             string       `json:"totalStaked"`
	ActiveNodeCount         int          `json:"activeNodeCount"`
	StandbyNodeCount        int          `json:"standbyNodeCount"`
	TotalReserve            string       `json:"totalReserve"`
	PoolShareFactor         string       `json:"poolShareFactor"`
	BlockRewards            *BlockReward `json:"blockRewards"`
	BondingRoi              string       `json:"bondingROI"`
	StakingRoi              string       `json:"stakingROI"`
	NextChurnHeight         string       `json:"nextChurnHeight"`
	PoolActivationCountdown int          `json:"poolActivationCountdown"`
}

type Node struct {
	Secp256k1 string `json:"secp256k1"`
	Ed25519   string `json:"ed25519"`
}

type Pool struct {
	Asset            string `json:"asset"`
	Status           string `json:"status"`
	Price            string `json:"price"`
	AssetStakedTotal string `json:"assetStakedTotal"`
	RuneStakedTotal  string `json:"runeStakedTotal"`
	PoolStakedTotal  string `json:"poolStakedTotal"`
	AssetDepth       string `json:"assetDepth"`
	RuneDepth        string `json:"runeDepth"`
	PoolDepth        string `json:"poolDepth"`
	PoolUnits        string `json:"poolUnits"`
	SellVolume       string `json:"sellVolume"`
	BuyVolume        string `json:"buyVolume"`
	PoolVolume       string `json:"poolVolume"`
	PoolVolume24hr   string `json:"poolVolume24hr"`
	SellTxAverage    string `json:"sellTxAverage"`
	BuyTxAverage     string `json:"buyTxAverage"`
	PoolTxAverage    string `json:"poolTxAverage"`
	SellSlipAverage  string `json:"sellSlipAverage"`
	BuySlipAverage   string `json:"buySlipAverage"`
	PoolSlipAverage  string `json:"poolSlipAverage"`
	SellFeeAverage   string `json:"sellFeeAverage"`
	BuyFeeAverage    string `json:"buyFeeAverage"`
	PoolFeeAverage   string `json:"poolFeeAverage"`
	SellFeesTotal    string `json:"sellFeesTotal"`
	BuyFeesTotal     string `json:"buyFeesTotal"`
	PoolFeesTotal    string `json:"poolFeesTotal"`
	SellAssetCount   string `json:"sellAssetCount"`
	BuyAssetCount    string `json:"buyAssetCount"`
	SwappingTxCount  string `json:"swappingTxCount"`
	SwappersCount    string `json:"swappersCount"`
	StakeTxCount     string `json:"stakeTxCount"`
	WithdrawTxCount  string `json:"withdrawTxCount"`
	StakingTxCount   string `json:"stakingTxCount"`
	StakersCount     string `json:"stakersCount"`
	AssetRoi         string `json:"assetROI"`
	RuneRoi          string `json:"runeROI"`
	PoolRoi          string `json:"poolROI"`
	PoolRoi12        string `json:"poolROI12"`
}

type PoolAddress struct {
	Chain   string `json:"chain"`
	PubKey  string `json:"pub_key"`
	Address string `json:"address"`
}

type PoolAddresses struct {
	Current []*PoolAddress `json:"current"`
}

type Staker struct {
	PoolsArray  []*string `json:"poolsArray"`
	TotalStaked string    `json:"totalStaked"`
	TotalEarned string    `json:"totalEarned"`
	TotalRoi    string    `json:"totalROI"`
}

type StakerPoolData struct {
	Asset            string `json:"asset"`
	StakeUnits       string `json:"stakeUnits"`
	DateFirstStaked  int    `json:"dateFirstStaked"`
	HeightLastStaked int    `json:"heightLastStaked"`
}

type StringValues struct {
	DefaultPoolStatus string `json:"DefaultPoolStatus"`
}

type Transaction struct {
	Pool    string               `json:"pool"`
	Type    string               `json:"type"`
	Status  string               `json:"status"`
	In      *TransactionDetail   `json:"in"`
	Out     []*TransactionDetail `json:"out"`
	Date    int                  `json:"date"`
	Gas     *AssetAmount         `json:"gas"`
	Options *TransactionOptions  `json:"options"`
	Height  string               `json:"height"`
	Events  *TransactionEvent    `json:"events"`
}

type TransactionDetail struct {
	TxID    string         `json:"txID"`
	Memo    string         `json:"memo"`
	Address string         `json:"address"`
	Coins   []*AssetAmount `json:"coins"`
}

type TransactionEvent struct {
	Fee        string `json:"fee"`
	StakeUnits string `json:"stakeUnits"`
	Slip       string `json:"slip"`
}

type TransactionOptions struct {
	PriceTarget         *string `json:"priceTarget"`
	WithdrawBasisPoints *string `json:"withdrawBasisPoints"`
	Asymmetry           *string `json:"asymmetry"`
}

type Transactions struct {
	Count int            `json:"count"`
	Txs   []*Transaction `json:"txs"`
}
