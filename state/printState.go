package state

import (
	"fmt"
	"sync"
)

var cntsMutex sync.Mutex
var cnts map[string]int

func PrintState(state *State) {
	if cnts == nil {
		cnts = make(map[string]int, 0)
	}

	str := fmt.Sprintf("%s %s", "===PrintState Start===\n", state.FactomNodeName)
	str = fmt.Sprintf("%s %s %s dbht: %d", str, "=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.\n", state.FactomNodeName, state.LLeaderHeight)
	str = fmt.Sprintf("%s =.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.\n", str)
	str = fmt.Sprintf("%s =.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.\n", str)
	str = fmt.Sprintf("%s =.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.\n", str)
	str = fmt.Sprintf("%s =.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.\n", str)
	str = fmt.Sprintf("%s =.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.=.\n", str)
	cntsMutex.Lock()
	str = fmt.Sprintf("%s %35s = %+v  Last value %d \n\n", str, "ResetCnt", state.ResetCnt, cnts[state.FactomNodeName])
	cnts[state.FactomNodeName] = state.ResetCnt
	cntsMutex.Unlock()

	str = fmt.Sprintf("%s %35s = %+v\n", str, "filename", state.filename)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Salt", state.Salt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Cfg", state.Cfg)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Prefix", state.Prefix)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FactomNodeName", state.FactomNodeName)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FactomdVersion", state.FactomdVersion)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LogPath", state.LogPath)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LdbPath", state.LdbPath)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "BoltDBPath", state.BoltDBPath)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LogLevel", state.LogLevel)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ConsoleLogLevel", state.ConsoleLogLevel)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "NodeMode", state.NodeMode)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBType", state.DBType)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "CloneDBType", state.CloneDBType)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ExportData", state.ExportData)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ExportDataSubpath", state.ExportDataSubpath)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LocalServerPrivKey", state.LocalServerPrivKey)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DirectoryBlockInSeconds", state.DirectoryBlockInSeconds)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "PortNumber", state.PortNumber)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DropRate", state.DropRate)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Delay", state.Delay)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ControlPanelPort", state.ControlPanelPort)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ControlPanelSetting", state.ControlPanelSetting)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ControlPanelChannel", state.ControlPanelChannel)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ControlPanelDataRequest", state.ControlPanelDataRequest)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Network", state.Network)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MainNetworkPort", state.MainNetworkPort)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "PeersFile", state.PeersFile)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MainSeedURL", state.MainSeedURL)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MainSpecialPeers", state.MainSpecialPeers)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "TestNetworkPort", state.TestNetworkPort)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "TestSeedURL", state.TestSeedURL)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "TestSpecialPeers", state.TestSpecialPeers)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LocalNetworkPort", state.LocalNetworkPort)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LocalSeedURL", state.LocalSeedURL)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LocalSpecialPeers", state.LocalSpecialPeers)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "CustomNetworkID", state.CustomNetworkID)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "IdentityChainID", state.IdentityChainID)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Identities", state.Identities)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Authorities", state.Authorities)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "AuthorityServerCount", state.AuthorityServerCount)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Status", state.Status)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "serverPrt", state.serverPrt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "starttime", state.starttime)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "transCnt", state.transCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "lasttime", state.lasttime)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "tps", state.tps)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBStateAskCnt", state.DBStateAskCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBStateAnsCnt", state.DBStateAnsCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBStateReplyCnt", state.DBStateReplyCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBStateFailsCnt", state.DBStateFailsCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MissingRequestSendCnt", state.MissingRequestSendCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MissingRequestReplyCnt", state.MissingRequestReplyCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MissingRequestIgnoreCnt", state.MissingRequestIgnoreCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MissingResponseAppliedCnt", state.MissingResponseAppliedCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ResendCnt", state.ResendCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ExpireCnt", state.ExpireCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "tickerQueue", state.tickerQueue)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "timerMsgQueue", state.timerMsgQueue)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "TimeOffset", state.TimeOffset)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MaxTimeOffset", state.MaxTimeOffset)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "networkOutMsgQueue", state.networkOutMsgQueue)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "networkInvalidMsgQueue", state.networkInvalidMsgQueue)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "inMsgQueue", state.inMsgQueue)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "apiQueue", state.apiQueue)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ackQueue", state.ackQueue)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "msgQueue", state.msgQueue)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ShutdownChan", state.ShutdownChan)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "JournalFile", state.JournalFile)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Journaling", state.Journaling)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "serverPrivKey", state.serverPrivKey)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "serverPubKey", state.serverPubKey)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "serverPendingPrivKeys", state.serverPendingPrivKeys)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "serverPendingPubKeys", state.serverPendingPubKeys)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "RpcUser", state.RpcUser)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "RpcPass", state.RpcPass)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "RpcAuthHash", state.RpcAuthHash)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FactomdTLSEnable", state.FactomdTLSEnable)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "factomdTLSKeyFile", state.factomdTLSKeyFile)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "factomdTLSCertFile", state.factomdTLSCertFile)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FactomdLocations", state.FactomdLocations)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "StartDelay", state.StartDelay)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "StartDelayLimit", state.StartDelayLimit)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "RunLeader", state.RunLeader)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LLeaderHeight", state.LLeaderHeight)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Leader", state.Leader)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LeaderVMIndex", state.LeaderVMIndex)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LeaderPL", state.LeaderPL)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "PLProcessHeight", state.PLProcessHeight)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "OneLeader", state.OneLeader)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "OutputAllowed", state.OutputAllowed)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "CurrentMinute", state.CurrentMinute)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EOMsyncing", state.EOMsyncing)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EOM", state.EOM)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EOMLimit", state.EOMLimit)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EOMProcessed", state.EOMProcessed)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EOMDone", state.EOMDone)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EOMMinute", state.EOMMinute)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EOMSys", state.EOMSys)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBSig", state.DBSig)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBSigLimit", state.DBSigLimit)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBSigProcessed", state.DBSigProcessed)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBSigDone", state.DBSigDone)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBSigSys", state.DBSigSys)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "KeepMismatch", state.KeepMismatch)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBSigFails", state.DBSigFails)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Newblk", state.Newblk)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Saving", state.Saving)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Syncing", state.Syncing)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "NetStateOff", state.NetStateOff)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DebugConsensus", state.DebugConsensus)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FactoidTrans", state.FactoidTrans)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ECCommits", state.ECCommits)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ECommits", state.ECommits)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FCTSubmits", state.FCTSubmits)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "NewEntryChains", state.NewEntryChains)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "NewEntries", state.NewEntries)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LeaderTimestamp", state.LeaderTimestamp)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "resendHolding", state.resendHolding)
	//str = fmt.Sprintf("%s %35s = %+v\n", str, "Holding", state.Holding)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "XReview", state.XReview)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Acks", state.Acks)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Commits", state.Commits)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "InvalidMessages", state.InvalidMessages)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "InvalidMessagesMutex", state.InvalidMessagesMutex)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "AuditHeartBeats", state.AuditHeartBeats)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FaultTimeout", state.FaultTimeout)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FaultWait", state.FaultWait)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EOMfaultIndex", state.EOMfaultIndex)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LastFaultAction", state.LastFaultAction)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LastTiebreak", state.LastTiebreak)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "AuthoritySetString", state.AuthoritySetString)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "NetworkNumber", state.NetworkNumber)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DB", state.DB)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Logger", state.Logger)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Anchor", state.Anchor)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "DBStates", state.DBStates)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ResetRequest", state.ResetRequest)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ProcessLists", state.ProcessLists)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "AuthorityDeltas", state.AuthorityDeltas)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FactoidState", state.FactoidState)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "NumTransactions", state.NumTransactions)
	//str = fmt.Sprintf("%s %35s = %+v\n", str, "FactoidBalancesP", state.FactoidBalancesP)
	//str = fmt.Sprintf("%s %35s = %+v\n", str, "FactoidBalancesPMutex", state.FactoidBalancesPMutex)
	//str = fmt.Sprintf("%s %35s = %+v\n", str, "ECBalancesP", state.ECBalancesP)
	//str = fmt.Sprintf("%s %35s = %+v\n", str, "ECBalancesPMutex", state.ECBalancesPMutex)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "Port", state.Port)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "IsReplaying", state.IsReplaying)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ReplayTimestamp", state.ReplayTimestamp)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MissingEntryBlockRepeat", state.MissingEntryBlockRepeat)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EntryBlockDBHeightComplete", state.EntryBlockDBHeightComplete)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EntryBlockDBHeightProcessing", state.EntryBlockDBHeightProcessing)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MissingEntryBlocks", state.MissingEntryBlocks)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MissingEntryRepeat", state.MissingEntryRepeat)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EntryDBHeightComplete", state.EntryDBHeightComplete)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EntryHeightComplete", state.EntryHeightComplete)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "EntryDBHeightProcessing", state.EntryDBHeightProcessing)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "MissingEntries", state.MissingEntries)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LastPrint", state.LastPrint)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "LastPrintCnt", state.LastPrintCnt)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FactoshisPerEC", state.FactoshisPerEC)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FERChainId", state.FERChainId)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "ExchangeRateAuthorityAddress", state.ExchangeRateAuthorityAddress)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FERChangeHeight", state.FERChangeHeight)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FERChangePrice", state.FERChangePrice)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FERPriority", state.FERPriority)
	str = fmt.Sprintf("%s %35s = %+v\n", str, "FERPrioritySetHeight", state.FERPrioritySetHeight)

	str = fmt.Sprintf("%s%s", str, "===PrintState End===")

	fmt.Println(str)
}
