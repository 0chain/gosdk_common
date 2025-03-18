package commonsdk

import (
	"time"

	"github.com/0chain/gosdk_common/core/common"
	"github.com/0chain/gosdk_common/zboxcore/blockchain"
	"github.com/0chain/gosdk_common/zboxcore/fileref"
)

// BlobberAllocationStats represents the blobber allocation statistics.
type BlobberAllocationStats struct {
	BlobberID        string
	BlobberURL       string
	ID               string `json:"ID"`
	Tx               string `json:"Tx"`
	TotalSize        int64  `json:"TotalSize"`
	UsedSize         int    `json:"UsedSize"`
	OwnerID          string `json:"OwnerID"`
	OwnerPublicKey   string `json:"OwnerPublicKey"`
	Expiration       int    `json:"Expiration"`
	AllocationRoot   string `json:"AllocationRoot"`
	BlobberSize      int    `json:"BlobberSize"`
	BlobberSizeUsed  int    `json:"BlobberSizeUsed"`
	LatestRedeemedWM string `json:"LatestRedeemedWM"`
	IsRedeemRequired bool   `json:"IsRedeemRequired"`
	CleanedUp        bool   `json:"CleanedUp"`
	Finalized        bool   `json:"Finalized"`
	Terms            []struct {
		ID           int    `json:"ID"`
		BlobberID    string `json:"BlobberID"`
		AllocationID string `json:"AllocationID"`
		ReadPrice    int    `json:"ReadPrice"`
		WritePrice   int    `json:"WritePrice"`
	} `json:"Terms"`
}

type ConsolidatedFileMetaByName struct {
	Name                string
	Type                string
	Path                string
	LookupHash          string
	Hash                string
	MimeType            string
	Size                int64
	NumBlocks           int64
	ActualFileSize      int64
	ActualNumBlocks     int64
	EncryptedKey        string
	FileMetaHash        string
	ThumbnailHash       string
	ActualThumbnailSize int64
	ActualThumbnailHash string
	Collaborators       []fileref.Collaborator
	CreatedAt           common.Timestamp
	UpdatedAt           common.Timestamp
}

// Terms represents Blobber terms. A Blobber can update its terms,
// but any existing offer will use terms of offer signing time.
type Terms struct {
	ReadPrice        common.Balance `json:"read_price"`  // tokens / GB
	WritePrice       common.Balance `json:"write_price"` // tokens / GB
	MaxOfferDuration time.Duration  `json:"max_offer_duration"`
}

// PriceRange represents a price range allowed by user to filter blobbers.
type PriceRange struct {
	Min uint64 `json:"min"`
	Max uint64 `json:"max"`
}

// IsValid price range.
func (pr *PriceRange) IsValid() bool {
	return pr.Min <= pr.Max
}

type AllocationStats struct {
	UsedSize                  int64  `json:"used_size"`
	NumWrites                 int64  `json:"num_of_writes"`
	NumReads                  int64  `json:"num_of_reads"`
	TotalChallenges           int64  `json:"total_challenges"`
	OpenChallenges            int64  `json:"num_open_challenges"`
	SuccessChallenges         int64  `json:"num_success_challenges"`
	FailedChallenges          int64  `json:"num_failed_challenges"`
	LastestClosedChallengeTxn string `json:"latest_closed_challenge"`
}

// BlobberAllocation represents the blobber in the context of an allocation
type BlobberAllocation struct {
	BlobberID       string         `json:"blobber_id"`
	Size            int64          `json:"size"`
	Terms           Terms          `json:"terms"`
	MinLockDemand   common.Balance `json:"min_lock_demand"`
	Spent           common.Balance `json:"spent"`
	Penalty         common.Balance `json:"penalty"`
	ReadReward      common.Balance `json:"read_reward"`
	Returned        common.Balance `json:"returned"`
	ChallengeReward common.Balance `json:"challenge_reward"`
	FinalReward     common.Balance `json:"final_reward"`
}

type Allocation struct {
	// ID is the unique identifier of the allocation.
	ID string `json:"id"`
	// Tx is the transaction hash of the latest transaction related to the allocation.
	Tx string `json:"tx"`

	// DataShards is the number of data shards.
	DataShards int `json:"data_shards"`

	// ParityShards is the number of parity shards.
	ParityShards int `json:"parity_shards"`

	// Size is the size of the allocation.
	Size int64 `json:"size"`

	// Expiration is the expiration date of the allocation.
	Expiration int64 `json:"expiration_date"`

	// Owner is the id of the owner of the allocation.
	Owner string `json:"owner_id"`

	// OwnerPublicKey is the public key of the owner of the allocation.
	OwnerPublicKey string `json:"owner_public_key"`

	// Payer is the id of the payer of the allocation.
	Payer string `json:"payer_id"`

	// Blobbers is the list of blobbers that store the data of the allocation.
	Blobbers []*blockchain.StorageNode `json:"blobbers"`

	// Stats contains the statistics of the allocation.
	Stats *AllocationStats `json:"stats"`

	// TimeUnit is the time unit of the allocation.
	TimeUnit time.Duration `json:"time_unit"`

	// WritePool is the write pool of the allocation.
	WritePool common.Balance `json:"write_pool"`

	// BlobberDetails contains real terms used for the allocation.
	// If the allocation has updated, then terms calculated using
	// weighted average values.
	BlobberDetails []*BlobberAllocation `json:"blobber_details"`

	// ReadPriceRange is requested reading prices range.
	ReadPriceRange PriceRange `json:"read_price_range"`

	// WritePriceRange is requested writing prices range.
	WritePriceRange PriceRange `json:"write_price_range"`

	// MinLockDemand is the minimum lock demand of the allocation.
	MinLockDemand float64 `json:"min_lock_demand"`

	// ChallengeCompletionTime is the time taken to complete a challenge.
	ChallengeCompletionTime time.Duration `json:"challenge_completion_time"`

	// StartTime is the start time of the allocation.
	StartTime common.Timestamp `json:"start_time"`

	// Finalized is the flag to indicate if the allocation is finalized.
	Finalized bool `json:"finalized,omitempty"`

	// Cancelled is the flag to indicate if the allocation is cancelled.
	Canceled bool `json:"canceled,omitempty"`

	// MovedToChallenge is the amount moved to challenge pool related to the allocation.
	MovedToChallenge common.Balance `json:"moved_to_challenge,omitempty"`

	// MovedBack is the amount moved back from the challenge pool related to the allocation.
	MovedBack common.Balance `json:"moved_back,omitempty"`

	// MovedToValidators is the amount moved to validators related to the allocation.
	MovedToValidators common.Balance `json:"moved_to_validators,omitempty"`

	// FileOptions is a bitmask of file options, which are the permissions of the allocation.
	FileOptions uint16 `json:"file_options"`

	IsEnterprise bool `json:"is_enterprise"`

	StorageVersion int `json:"storage_version"`

	// Owner ecdsa public key
	OwnerSigningPublicKey string `json:"owner_signing_public_key"`

	// FileOptions to define file restrictions on an allocation for third-parties
	// default 00000000 for all crud operations suggesting only owner has the below listed abilities.
	// enabling option/s allows any third party to perform certain ops
	// 		00000001 - 1  - upload
	// 		00000010 - 2  - delete
	// 		00000100 - 4  - update
	// 		00001000 - 8  - move
	// 		00010000 - 16 - copy
	// 		00100000 - 32 - rename
	ThirdPartyExtendable bool `json:"third_party_extendable"`
}

// UpdateTerms represents Blobber terms during update blobber calls.
// A Blobber can update its terms, but any existing offer will use terms of offer signing time.
type UpdateTerms struct {
	ReadPrice        *common.Balance `json:"read_price,omitempty"`  // tokens / GB
	WritePrice       *common.Balance `json:"write_price,omitempty"` // tokens / GB
	MaxOfferDuration *time.Duration  `json:"max_offer_duration,omitempty"`
}
