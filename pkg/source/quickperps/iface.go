package quickperps

// NOTE: we generate mock files in the same package to avoid cycle dependencies
//go:generate mockgen -destination ./mock_vault_reader.go -package gmx pool-service/internal/dex/gmx IVaultReader
//go:generate mockgen -destination ./mock_vault_price_feed_reader.go -package gmx pool-service/internal/dex/gmx IVaultPriceFeedReader
//go:generate mockgen -destination ./mock_fast_price_feed_reader_v1.go -package gmx pool-service/internal/dex/gmx IFastPriceFeedV1Reader
//go:generate mockgen -destination ./mock_price_feed_reader.go -package gmx pool-service/internal/dex/gmx IPriceFeedReader
//go:generate mockgen -destination ./mock_usdq_reader.go -package gmx pool-service/internal/dex/gmx IUSDQReader
//go:generate mockgen -destination ./mock_chainlink_flags_reader.go -package gmx pool-service/internal/dex/gmx IChainlinkFlagsReader
//go:generate mockgen -destination ./mock_pancake_pair_reader.go -package gmx pool-service/internal/dex/gmx IPancakePairReader

import (
	"context"
	"math/big"
)

// IVaultReader reads vault smart contract
type IVaultReader interface {
	Read(ctx context.Context, address string) (*Vault, error)
}

// IVaultPriceFeedReader reads vault price feed smart contract
type IVaultPriceFeedReader interface {
	Read(ctx context.Context, address string, tokens []string) (*VaultPriceFeed, error)
}

// IFastPriceFeedV1Reader reads fast price feed smart contract
type IFastPriceFeedV1Reader interface {
	Read(ctx context.Context, address string, tokens []string) (*FastPriceFeedV1, error)
}

type IFastPriceFeedV2Reader interface {
	Read(ctx context.Context, address string, tokens []string) (*FastPriceFeedV2, error)
}

// IPriceFeedReader reads price feed smart contract
type IPriceFeedReader interface {
	Read(ctx context.Context, address string) (*PriceFeed, error)
}

// IUSDQReader reads usdq smart contract
type IUSDQReader interface {
	Read(ctx context.Context, address string) (*USDQ, error)
}

// IChainlinkFlagsReader reads chainlink flag smart contract
type IChainlinkFlagsReader interface {
	Read(ctx context.Context, address string) (*ChainlinkFlags, error)
}

// IPancakePairReader reads pancake pair smart contract
type IPancakePairReader interface {
	Read(ctx context.Context, address string) (*PancakePair, error)
}

type IFastPriceFeed interface {
	GetVersion() int
	GetPrice(token string, refPrice *big.Int, maximise bool) *big.Int
}
