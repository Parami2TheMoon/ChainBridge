// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	events "github.com/parami-protocol/chainbridge-substrate-events"
)

type EventErc721Minted struct {
	Phase   types.Phase
	Owner   types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Transferred struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Burned struct {
	Phase   types.Phase
	TokenId types.AccountID
	Topics  []types.Hash
}

type EventXAssetsRemark struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// EventNFTDeposited is emitted when NFT is ready to be deposited to other chain.
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

// EventFeeChanged is emitted when a fee for a given key is changed.
type EventFeeChanged struct {
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
}

// EventNewMultiAccount is emitted when a multi account has been created.
// First param is the account that created it, second is the multisig account.
type EventNewMultiAccount struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultiAccountUpdated is emitted when a multi account has been updated. First param is the multisig account.
type EventMultiAccountUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventMultiAccountRemoved is emitted when a multi account has been removed. First param is the multisig account.
type EventMultiAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventNewMultisig is emitted when a new multisig operation has begun.
// First param is the account that is approving, second is the multisig account.
type EventNewMultisig struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultisigApproval is emitted when a multisig operation has been approved by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigApproval struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventMultisigExecuted is emitted when a multisig operation has been executed by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigExecuted struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Result    types.DispatchResult
	Topics    []types.Hash
}

// EventMultisigCancelled is emitted when a multisig operation has been cancelled by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigCancelled struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

type EventTreasuryMinting struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventRadClaimsClaimed is emitted when RAD Tokens have been claimed
type EventRadClaimsClaimed struct {
	Phase  types.Phase
	Who    types.AccountID
	Value  types.U128
	Topics []types.Hash
}

// EventRadClaimsRootHashStored is emitted when RootHash has been stored for the correspondent RAD Claims batch
type EventRadClaimsRootHashStored struct {
	Phase    types.Phase
	RootHash types.Hash
	Topics   []types.Hash
}

// EventNftTransferred is emitted when the ownership of the asset has been transferred to the account
type EventNftTransferred struct {
	Phase      types.Phase
	RegistryId RegistryId
	AssetId    AssetId
	Who        types.AccountID
	Topics     []types.Hash
}

// EventRegistryMint is emitted when successfully minting an NFT
type EventRegistryMint struct {
	Phase      types.Phase
	RegistryId RegistryId
	TokenId    TokenId
	Topics     []types.Hash
}

// EventRegistryRegistryCreated is emitted when successfully creating a NFT registry
type EventRegistryRegistryCreated struct {
	Phase      types.Phase
	RegistryId RegistryId
	Topics     []types.Hash
}

// EventRegistryTmp is emitted only for testing
type EventRegistryTmp struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

type EventBalancesWithdraw struct {
	Phase   types.Phase
	Who     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

type EventElectionProviderMultiPhaseSolutionStored struct {
	Phase           types.Phase
	ElectionCompute types.ElectionCompute
	PreviousEjected bool
	Topics          []types.Hash
}
type EventElectionProviderMultiPhaseElectionFinalized struct {
	Phase           types.Phase
	ElectionCompute OptionElectionCompute
	Topics          []types.Hash
}

type EventElectionProviderMultiPhaseRewarded struct {
	Phase     types.Phase
	AccountID types.AccountID
	Value     types.U128
	Topics    []types.Hash
}

type EventElectionProviderMultiPhaseSlashed struct {
	Phase     types.Phase
	AccountID types.AccountID
	Value     types.U128
	Topics    []types.Hash
}

type EventElectionProviderMultiPhaseSignedPhaseStarted struct {
	Phase  types.Phase
	U32    types.U32
	Topics []types.Hash
}

type EventElectionProviderMultiPhaseUnsignedPhaseStarted struct {
	Phase  types.Phase
	U32    types.U32
	Topics []types.Hash
}

type EventStakingEraPaid struct {
	Phase     types.Phase
	Index     types.U32
	Payout    types.U128
	Remainder types.U128
	Topics    []types.Hash
}

type EventStakingStakersElected struct {
	Phase  types.Phase
	Topics []types.Hash
}

type EventUtilityItemCompleted struct {
	Phase  types.Phase
	Topics []types.Hash
}

type EventAdCreated struct {
	Phase   types.Phase
	ID      types.H256
	Creator types.H160
	Topics  []types.Hash
}

type EventAdDeposited struct {
	Phase  types.Phase
	ID     types.H256
	Owner  types.H160
	Value  types.U128
	Topics []types.Hash
}

type EventAdUpdated struct {
	Phase  types.Phase
	ID     types.H256
	Topics []types.Hash
}

type EventAdBid struct {
	Phase  types.Phase
	KOL    types.H160
	ID     types.H256
	Value  types.U128
	Topics []types.Hash
}

type EventAdEnd struct {
	Phase  types.Phase
	KOL    types.H160
	ID     types.H256
	Value  types.U128
	Topics []types.Hash
}

type EventAdPaid struct {
	Phase   types.Phase
	ID      types.H256
	AssetID types.U32
	Visitor types.H160
	Reward  types.U128
	Referer types.OptionH160
	Award   types.U128
	Topics  []types.Hash
}

type EventAdSwapTriggered struct {
	Phase  types.Phase
	ID     types.H256
	KOL    types.H160
	Remain types.U128
	Topics []types.Hash
}

type EventAdvertiserDeposited struct {
	Phase  types.Phase
	DID    types.H160
	Value  types.U128
	Topics []types.Hash
}

type EventAdvertiserBlocked struct {
	Phase  types.Phase
	DID    types.H160
	Topics []types.Hash
}

type EventDidAssigned struct {
	Phase     types.Phase
	DID       types.H160
	AccountID types.AccountID
	Referrer  types.OptionH160
	Topics    []types.Hash
}

type EventDidRevoked struct {
	Phase  types.Phase
	DID    types.H160
	Topics []types.Hash
}

type EventDidTransferred struct {
	Phase  types.Phase
	DID    types.H160
	From   types.AccountID
	To     types.AccountID
	Topics []types.Hash
}

type EventDidUpdated struct {
	Phase  types.Phase
	DID    types.H160
	Topics []types.Hash
}

type EventLinkerAccountLinked struct {
	Phase   types.Phase
	DID     types.H160
	Type    types.U8
	Account types.Bytes
	By      types.H160
	Topics  []types.Hash
}

type EventLinkerAccountUnlinked struct {
	Phase  types.Phase
	DID    types.H160
	Type   types.U8
	Topics []types.Hash
}

type EventLinkerBlocked struct {
	Phase  types.Phase
	DID    types.H160
	Topics []types.Hash
}

type EventLinkerDeposited struct {
	Phase  types.Phase
	DID    types.H160
	Value  types.U128
	Topics []types.Hash
}

type EventLinkerTrusted struct {
	Phase  types.Phase
	DID    types.H160
	Topics []types.Hash
}

type EventLinkerValidationFailed struct {
	Phase   types.Phase
	DID     types.H160
	Type    types.U8
	Account types.Bytes
	Topics  []types.Hash
}

type EventMagicCreated struct {
	Phase      types.Phase
	Stash      types.AccountID
	Controller types.AccountID
	Topics     []types.Hash
}

type EventMagicChanged struct {
	Phase      types.Phase
	Stash      types.AccountID
	Controller types.AccountID
	Topics     []types.Hash
}

type EventMagicCodo struct {
	Phase  types.Phase
	Result types.DispatchResult
	Topics []types.Hash
}

type EventNftBacked struct {
	Phase  types.Phase
	DID    types.H160
	KOL    types.H160
	Value  types.U128
	Topics []types.Hash
}

type EventNftClaimed struct {
	Phase  types.Phase
	DID    types.H160
	KOL    types.H160
	Value  types.U128
	Topics []types.Hash
}

type EventNftMinted struct {
	Phase    types.Phase
	KOL      types.H160
	Class    types.U32
	Instance types.U32
	Name     types.Bytes
	Symbol   types.Bytes
	Value    types.U128
	Topics   []types.Hash
}

type EventSwapCreated struct {
	Phase   types.Phase
	TokenID types.U32
	Topics  []types.Hash
}

type EventSwapLiquidityAdded struct {
	Phase     types.Phase
	ID        types.U32
	AccountID types.AccountID
	Currency  types.U128
	Tokens    types.U128
	Topics    []types.Hash
}

type EventSwapLiquidityRemoved struct {
	Phase     types.Phase
	ID        types.U32
	AccountID types.AccountID
	Currency  types.U128
	Tokens    types.U128
	Topics    []types.Hash
}

type EventSwapTokenBought struct {
	Phase     types.Phase
	ID        types.U32
	AccountID types.AccountID
	Tokens    types.U128
	Currency  types.U128
	Topics    []types.Hash
}

type EventSwapTokenSold struct {
	Phase     types.Phase
	ID        types.U32
	AccountID types.AccountID
	Tokens    types.U128
	Currency  types.U128
	Topics    []types.Hash
}

type EventTagCreated struct {
	Phase   types.Phase
	Hash    types.Bytes
	Creator types.H160
	Topics  []types.Hash
}

type Events struct {
	types.EventRecords
	events.Events

	Balances_Withdraw                               []EventBalancesWithdraw                               //nolint:stylecheck,golint
	ElectionProviderMultiPhase_SolutionStored       []EventElectionProviderMultiPhaseSolutionStored       //nolint:stylecheck,golint
	ElectionProviderMultiPhase_ElectionFinalized    []EventElectionProviderMultiPhaseElectionFinalized    //nolint:stylecheck,golint
	ElectionProviderMultiPhase_Rewarded             []EventElectionProviderMultiPhaseRewarded             //nolint:stylecheck,golint
	ElectionProviderMultiPhase_Slashed              []EventElectionProviderMultiPhaseSlashed              //nolint:stylecheck,golint
	ElectionProviderMultiPhase_SignedPhaseStarted   []EventElectionProviderMultiPhaseSignedPhaseStarted   //nolint:stylecheck,golint
	ElectionProviderMultiPhase_UnsignedPhaseStarted []EventElectionProviderMultiPhaseUnsignedPhaseStarted //nolint:stylecheck,golint
	Staking_StakersElected                          []EventStakingStakersElected                          //nolint:stylecheck,golint
	Staking_EraPaid                                 []EventStakingEraPaid                                 //nolint:stylecheck,golint

	Erc721_Minted                    []EventErc721Minted                   //nolint:stylecheck,golint
	Erc721_Transferred               []EventErc721Transferred              //nolint:stylecheck,golint
	Erc721_Burned                    []EventErc721Burned                   //nolint:stylecheck,golint
	Nfts_DepositAsset                []EventNFTDeposited                   //nolint:stylecheck,golint
	Council_Proposed                 []types.EventCollectiveProposed       //nolint:stylecheck,golint
	Council_Voted                    []types.EventCollectiveVoted          //nolint:stylecheck,golint
	Council_Approved                 []types.EventCollectiveApproved       //nolint:stylecheck,golint
	Council_Disapproved              []types.EventCollectiveDisapproved    //nolint:stylecheck,golint
	Council_Executed                 []types.EventCollectiveExecuted       //nolint:stylecheck,golint
	Council_MemberExecuted           []types.EventCollectiveMemberExecuted //nolint:stylecheck,golint
	Council_Closed                   []types.EventCollectiveClosed         //nolint:stylecheck,golint
	Fees_FeeChanged                  []EventFeeChanged                     //nolint:stylecheck,golint
	MultiAccount_NewMultiAccount     []EventNewMultiAccount                //nolint:stylecheck,golint
	MultiAccount_MultiAccountUpdated []EventMultiAccountUpdated            //nolint:stylecheck,golint
	MultiAccount_MultiAccountRemoved []EventMultiAccountRemoved            //nolint:stylecheck,golint
	MultiAccount_NewMultisig         []EventNewMultisig                    //nolint:stylecheck,golint
	MultiAccount_MultisigApproval    []EventMultisigApproval               //nolint:stylecheck,golint
	MultiAccount_MultisigExecuted    []EventMultisigExecuted               //nolint:stylecheck,golint
	MultiAccount_MultisigCancelled   []EventMultisigCancelled              //nolint:stylecheck,golint
	TreasuryReward_TreasuryMinting   []EventTreasuryMinting                //nolint:stylecheck,golint
	Nft_Transferred                  []EventNftTransferred                 //nolint:stylecheck,golint
	RadClaims_Claimed                []EventRadClaimsClaimed               //nolint:stylecheck,golint
	RadClaims_RootHashStored         []EventRadClaimsRootHashStored        //nolint:stylecheck,golint
	Registry_Mint                    []EventRegistryMint                   //nolint:stylecheck,golint
	Registry_RegistryCreated         []EventRegistryRegistryCreated        //nolint:stylecheck,golint
	Registry_RegistryTmp             []EventRegistryTmp                    //nolint:stylecheck,golint
	Utility_ItemCompleted            []EventUtilityItemCompleted           //nolint:stylecheck,golint

	Ad_Created              []EventAdCreated              //nolint:stylecheck,golint
	Ad_Deposited            []EventAdDeposited            //nolint:stylecheck,golint
	Ad_Updated              []EventAdUpdated              //nolint:stylecheck,golint
	Ad_Bid                  []EventAdBid                  //nolint:stylecheck,golint
	Ad_End                  []EventAdEnd                  //nolint:stylecheck,golint
	Ad_Paid                 []EventAdPaid                 //nolint:stylecheck,golint
	Ad_SwapTriggered        []EventAdSwapTriggered        //nolint:stylecheck,golint
	Advertiser_Deposited    []EventAdvertiserDeposited    //nolint:stylecheck,golint
	Advertiser_Blocked      []EventAdvertiserBlocked      //nolint:stylecheck,golint
	Did_Assigned            []EventDidAssigned            //nolint:stylecheck,golint
	Did_Revoked             []EventDidRevoked             //nolint:stylecheck,golint
	Did_Transferred         []EventDidTransferred         //nolint:stylecheck,golint
	Did_Updated             []EventDidUpdated             //nolint:stylecheck,golint
	Linker_AccountLinked    []EventLinkerAccountLinked    //nolint:stylecheck,golint
	Linker_AccountUnlinked  []EventLinkerAccountUnlinked  //nolint:stylecheck,golint
	Linker_Blocked          []EventLinkerBlocked          //nolint:stylecheck,golint
	Linker_Deposited        []EventLinkerDeposited        //nolint:stylecheck,golint
	Linker_Trusted          []EventLinkerTrusted          //nolint:stylecheck,golint
	Linker_ValidationFailed []EventLinkerValidationFailed //nolint:stylecheck,golint
	Magic_Created           []EventMagicCreated           //nolint:stylecheck,golint
	Magic_Changed           []EventMagicChanged           //nolint:stylecheck,golint
	Magic_Codo              []EventMagicCodo              //nolint:stylecheck,golint
	Nft_Backed              []EventNftBacked              //nolint:stylecheck,golint
	Nft_Claimed             []EventNftClaimed             //nolint:stylecheck,golint
	Nft_Minted              []EventNftMinted              //nolint:stylecheck,golint
	Swap_Created            []EventSwapCreated            //nolint:stylecheck,golint
	Swap_LiquidityAdded     []EventSwapLiquidityAdded     //nolint:stylecheck,golint
	Swap_LiquidityRemoved   []EventSwapLiquidityRemoved   //nolint:stylecheck,golint
	Swap_SwapBuy            []EventSwapTokenBought        //nolint:stylecheck,golint
	Swap_SwapSell           []EventSwapTokenSold          //nolint:stylecheck,golint
	Tag_Created             []EventTagCreated             //nolint:stylecheck,golint
	XAssets_Remark          []EventXAssetsRemark          //nolint:stylecheck,golint
}
