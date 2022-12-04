package main
import(
	"fmt"
	"strings"
)
type Rucksack struct{
	Com1 string
	Com2 string
}
type RucksackDups struct{
	Com1 string
	Com2 string
	Dup rune
}
type RucksackDupsValue struct{
	Com1 string
	Com2 string
	Dup rune
	value int
}



func main(){

	input := `hDsDDttbhsmshNNWMNWGbTNqZq
	VQfjnlFvnQFRdZWdVtqMGdWW
	zvvvRnFFfjjlRBlBPzgQgRvvmtrmhHcptLHCDhcHHmLsBmsB
	FrzFvvdTDcTnmTzdDTTzdvWmjhgVPrhSljSQSPwPjPjPjSVC
	sMsGbqGsbbRqRbBMBGRMbLpNSSpjhlQljHVClhjgPjjPhlVp
	sNbGtJbMfssNtvcnWFVmnvDd
	TNfmdFJmfdZMQffVRQVV
	jVHBCcDSjWrMZjvg
	SShSbCGpcBtBtwtVLJJddmtLmT
	CtpNftbNWbtSJDHqGZJFLfLr
	dPsHlsRBHcZdqDFDZwwJ
	snjVlvTPlPjVlQlHWjpSmzgNNzSmtpSm
	qhZtSVqCqThGcGzZnnfZcB
	WbddWbDwrBzcpzHpBb
	DBBMFWRJDrDFWLWljCqjQjFvtCsqTjqs
	vhFTzRzzTmPvbplWFtQttQQZtZhMZqcqSQ
	fJVCfDfJNCLDwJNGmssZgwqgZcmtgcms
	VmdVNLHGGVDBdfLCHnLGHnbWpTplWbddRTlzplWPpbFp
	smwtNVqRjNmZjZBDSvzSzl
	FnTJFcTTFccCrJGTLncdCCcPJZfBBDSlSJwZDlggSffvgSSf
	FWTFGLFWLCPWrCnFnQWNbQVVwphhmHHbptVsss
	BrrgrtgfBpPFhhgMWq
	ZGvsvDGClvsSRScpGBhPphWMPhhTNh
	SBSBdBZCcdwHrVQwQr
	qmFqdVtqsVdzqGbwMJwGPpmPHM
	ZjTjLQLLDrrLjcFhlfrGHppfbJwGpMMpGHwRCC
	BTlZjrBBBcLTcDrjBlThWBjBtWNqnSzSnsSdvsFggNnqVVnv
	zVvmjGgpcJnbTTTJHRHSRb
	NPFrFQfCLPrdRlbtQRRvBtHb
	frMLqPFMrfrPLCwqqNvjczwwnmwggGmssnmnnm
	HbFJhhshsffcvslmGmLFrQBrlFTG
	jNRPwwPjSPCdCdvRzRBTlRzmGrmB
	nCSWdNjCCPqvtttZnDscth
	ScSrRTPcSSDRWSptWcdmmWGbmGGLmLJvNJNbbJ
	flzHjFpZjFfjjgszjlqzJNnnvsmbMmGvJLNNbmJv
	jFZFVpffpVlqfQhgtTwcTVtrTcRBwBTD
	rHrdGSMSSbZbjShj
	qZfDBBvllvvWLtqbbQwhJjbtwbnQgN
	WzWlzZmLWBLZLzCzrHMVcRrMRFRCCccF
	BzdplppDlBBrqWnjFMBWqNWq
	whZhZSSHhhVSrvSgHPvgvjnFTPsFFnnNTcjTFnTTsc
	HLCwVSfZLffSHhLvwtQbJbrdGlRRrGdmpztG
	grDFfDlfCftCzCfNztclNFrBNQjbZjJjjPPVsjNvsbvPsj
	HwwTpGpRwMdpHWhvjzVsPJjJGVvVZs
	ShphRpwzMWdpHdwWHwMpnHwLDmmgcLCCDtLDCCSlcgcfCD
	ccqqLLqCqTSlZMLQMllZTvnNfjddttmmDpRJjvhfpfthRmdf
	rbVssWwggFrGsWzPbVFGJpftQRPDmQDpdPmdJmfR
	bBHWWHzWHWWHrsVbsFgwbFqqMLTnBclTMMSnLnqLqMQZ
	CcSPGCCPrdPtdjcsBLDghbVLhqDl
	vMJwTHzPvzVwqBBDblls
	QNvMFfRMJHPZHjnfWmdftSjSnp
	dnBCPhhBCrQfChdbNVGLszzDzVDsTbWT
	HgcJgpppPqqHwPwJSczVzDssNNVWTtqVVGts
	JHRFpjFccplcRwPJpHPScpMPQmCdBQQfQjjhCfBCrQQvdmnn
	rQGmVRLRbDRHmmZLGBGVLHBVFspSstWWWNJcsgpQTSsNppJS
	qlldhPdfCgnspJFWCFsFNT
	jMMzwndfhnwPfqPMjgjMVBrGBmrHDHLZbjRGHbHj
	wzpZfzHRSRfzgHfffZwwStCtSrBhBBCTrtFhhBFG
	QPjQQQDcDWJNFWtrtWrGmTCMtmBW
	lDvvDQcdjQcQLvlDwnpFgbbznZZglZsl
	RfMFTMFrVrSRFPlFSfVlHpLqgzpHBLzHBBVzVpHG
	CchbhcwdmdJmwJJtGgnqzppLmGGBQqQp
	hCsstCJwLvMRvsZsTR
	fQlfMlNClQhhZhrlWrWw
	njDbnTDTBtGjmrGvSh
	bgshBdBcDbTTdnnnTqcqLgqfpfQppCsCMsHHVVpHHNCHFF
	PbCnTbzJnqQNzbbTNDdpwcmjDmwjGQjccw
	hWgvSdLvwcGjSpSm
	vVfrFvvhHFTZndJq
	FFvRVCRqVRcfsDLrgqGNWjjHfhQQzGWjQHzN
	ppJPBwplwSBJTmPpTzWWStzHHjNNLNzHNh
	wnwMPbLJMJllJJwBmJmnLVvCvsCsbFgDrRrsCsvrqc
	jqHgVgdgGQttWCtNqNflmllgFnfDnmFFlpcl
	TZZsrrwwhwrsrZRGmhcfSnhGlmSFcf
	rsLvvbJPPLBGPCQqHWtMCMHN
	WzzBpCBpMsBpCvCfsgnPPfHgbfFNfF
	jtdTLLjGTGjDjLbbbGlDLLLmfFgmmfgrPrgmNPHSnnNnqFSq
	jbRRbjlwlRVpvWwvzBpB
	qpwzCzCznFznTcCvrcrvVcLb
	cPmNMHSlMsLfvWgsrWvL
	mMHGPDMBGSGPHlPBcBPzznnzpdQFBjqRFdwdnn
	QGZLJzmJrZgZzZhNQFqDWlWPWDFCWNRlPW
	hMMbhVbhHWCsPWCMRs
	BwjbSHVBVvfcTfZgZzQhrzdGrt
	cvPTjfDPpDmmBjbQjZMdlBZj
	CHNnghNChVzNgrFVwCMJLMMMMMQQdbLZ
	FNSzShrHhNnWgVnWfvfbpfpDGTfvsG
	FFpVrZhpTlSQlQzTtRtZHfmPmJDbRtZJ
	jNnwBLnWwBgNBNCwsNgsMsCLVfDWfmDJRffmRRtmfRmPDRtR
	BnLdNdjnBNcLngdCVBndgllFlQrrpQlSSFrQphFcQq
	JmVLJPMNjmVJpMLJSVmNQZZQZZrnTHqZQHTrTTMr
	ltfdwChhwRdRswDDdnBQqqWNTqrrHrqZRr
	dwdwwDwsGlhhDFtsCwhsJLcPmGPcGzSjSjjLzNPS
	QgSbgQCLQSJFMccLFLVVzH
	WBNffrBpBNdNRdWDfptBtdzWcMZZVPMVwHMmsVHFccHsDsVc
	ptphRWrfGRGRnqlCSQvhqbCzJS
	VvdMLMLMBMlVlVschsNpDGpdNsGc
	tqFSmnmnnttGfDqNcfvNDD
	nzRHnrwrrWRrrzHtbMlTMBjCvWLgBBMl
	pCBlRvzwzlCzvZqqDwzmvgtsLsQdgZsgPNtdrsWrst
	JbGjbGVGHSFbhbnhTShSTbQtQrPsLsHLQgNRtsgdQsLP
	GnnbJbGMGbSjjbSbCzqlwMRRDzqzMBBv
	TTVRJVMWMshSQtjSVTQJRQlcCBncJccdppnJcBBDngFpgP
	frvfzfHrwzZNrtNwzzzZncCCZFdFFCgBDpPcBcBg
	LwNGbqwvmvvtbNQjQlVVRWTGTQWM
	RnggwVLRLDfCVZhfpDGGMGMGcGzGNHvv
	jmmWBTSsBmFmSzctsqpccHvzpN
	SBSrTblSQPbQhNwwZfPVdZPf
	CCvCwzfNStLzfrbmMJbZMtlsbJMW
	gPPPBqDjBcPFpVgBRbnMsVsbJZnWsdbSbM
	qgDPHjHcPhpDRRpPBRpCGSLwvHQwGfzrHLGLTL
	CLGqDZZLTdddPsdJpq
	gbRbbnghnrWvgrdJdSTRSsVNJlld
	hMnwrjnnjggvnLDwGffTfwCZZZ
	NzJHbNHNNzJzgmHmzpQSvvLqbLsVVsVGvB
	WtWhtWDdrZldDWrWTlZgppVVsqQTVQBqsGqBsQVp
	jWjWRRRlPcHRwJgw
	CCnnFTmnPCMCRNfnwGwdfzvwwl
	VQQVShDSSshhDDtDLhjccGjLBBzBzlZflNZvwZzdwBzpSNNZ
	QLHhJDDhDhgscgtjbGHTrWbrTbRmbFrm
	CJbLvJvbwtFHqvLzwJqqqtHWTWRgDScDRSWQQjTRcWRDLT
	mssGsMNphZMNsPPBnhSjRgdnQRdWgdjgrcDn
	GGmBMMsffmslMGshZlMphGqCzHCbzlbvzqwzzgFJggCF
	WCgWBphpWLQZQpgdhGdwmfbfFRVRjRTbbSFttdbSbT
	qqrZnDNqZJDTVzRjVFbfSN
	rPMvqJJqrJMPJnZMZZgLPgLWLggWQghwhmBh
	CWGGzdHHmPPSmPsC
	LqwlZwRLrPMQlMqrlbZrQRsSNsmssSNSsNcBNpgmsJ
	lwLDQhrDMQqPGfhzGGjhGn
	ZqDlZssCqJJMvpdBpBBmBQSMRp
	wLgVcbgFLzTLTNNZmNNRdjdRmF
	HZHbctWTwgVWgsfrnnPqlnsWlD
	RSnwSPFcLnFPnRwjzctzbGNlZgNbbGdGpLhZdpgM
	BqqBfMvTmmJqDgGNVdGVbVJJZG
	vsTDfqBmHmMWQCwjtrHjjSFFnRSn
	LsCmmcDHRjdtNMstwwzJ
	TvThqfBFBNTnnndTtL
	lvGQfbFQGblFrRccLRSSlPPVHS
	qbLpqTHSqpbqbrPcQgjPDjcdDL
	gnzhhBBwBWZzMglmjDrDPjvfdvQPdwtd
	WZZzZZlZmhsMmFgRBBBzHHJpNJGVRSJVGTVpNqCH
	JDphhGhDdGzWRBnvqqLDNLMnLw
	gsrTHHffTHPcrPrlHCNZhvBZnZNLhPvBNwvh
	sHVCSsSghJpSJjQh
	JTMGlfjlTdqjnqbnqFwqmnbQ
	PBZhBBcWRZprPZcZDDCZTZRgnnzwbbsbnvhbvznFsNFFvvNQ
	BcWrVgCCZTDRDrlGttfVtMddSJlH
	vwwvpVbSvnSRRmfMCmTHVHTBHB
	QLZgDPgSDgGTMZfmTTBZ
	QDDsQFDlzlgtJlLdFDgSJQFvvpRvjqzjRwwwzWvhvWwqjj
	mRRTGGNNflGRGGmmgRblsGwCZwVZlZjVwjztpjZhpBCB
	PMLLFLHPLPnLqDDLvFDrzzMjhwVCjtphBzjMhV
	FDdSPSpLcDsNRRWSTNWN
	STldJthdJbtTqljCRDDHmqmj
	VVvNwwvNFssJFJPNNwVvRMCgCgDqjjjqrDqqMHqP
	QBZwQfZwfVhtcSBtBJnT
	TzjjPzsQTslNlNzPRVGJJJGGtTJmgJHtmTZC
	dBDWScMBhhPGgdwwJPfw
	SqSqbSPDBhqnMqvrrSWVNFpRVRLzVQslvpNjVL
	bWFgFCPFtgvDZWgtChDNFJHvGVzHHpjzHnnzGzzHRR
	qcScQbbmqdQmlQmrlcQwLmHlRRjzGHnHJnnGVjHzBHzG
	TqQwmLmfcddfwrfCgbWCNPsZNfCb
	pddprrtrCPdvJdMjwwwHnLwwjLWCLg
	qhzZTmZcmRhmpFlVHcQQVwWQHVQwnH
	lGmhfRfmBZRlmmbvDPBMvNbvJJpP
	NsptgfGLLNwnNQSZbCvZnRnMCb
	JldhdzwzBMCSZvrz
	JFcdWTdwhPTFVDVmTJNqmstLqgLtLtjGGpsG
	dVVTSgTDpHVDjgdWpdpHTZSbWGrnnvrNwzFGNrFwnNNwvh
	CPRlMPJcMQcBcsmmLCMPrzbFfhwfrvLrNNwwGwfF
	CRJmtbmJlQbsQlRBpZDVjTHTdjDtSjZt
	rQVJrRFdrwDfzHQHQBTnpWTW
	PCLbPcPCsgqCgPgLjScSqNbHTzMtWmWtzlTHmBtTlMssMT
	cCqghSSPcvgScPbwFGdDDVZFfDhZGB
	zrRQRdqzPHQtnMPrtzPMRRQMVBBblJJBSClBpJbpdCCbBlCC
	hTcGwzswGwGmGfDvvfGmGNfBpllVSWbWppNCBNBVpCBClW
	gvGFTmTgwDhTDccsTfzfmfGGQPgPPqrgRZPHnRqRZrQLnRgP
	hvmmJllPbmCRMNGMMlNwNl
	PFTpTVjTgpTpBRgMGMnRNHBB
	WWrqzTTPVQDPqpjTqPJbmLtcfsQsftbLbvct
	SzrmpjjcsjTZNzgnnNzN
	BLHNDwBLBPLwLBhwDVLgdQCgCQGTngHQZCngZd
	PPJBDvBVVBmppNjJjrrr
	ZHBNQFhsqHBsgCfqtctcPvSwPqrV
	LlnGTnJpJJTmdDpmLlmLndWfVrPvvRwDfcwwwRVwcfQtvP
	GnbQblWmWGdTJQdTGnZHsHhZhFNsbCsjFgjC
	hWfDzDTVndDMhddMlBWMBDfJRnRtvvSSQjCvZCtjtpJvSR
	bGHsccFcbscsqGPHNGcrpjJZtvSRtFtQCZrjSj
	GsbwGGwNNGLgPLwMzBzfMMVMTLdTCC
	GBcNzTSSmGzmTLNgvwgpNCDqpDggpw
	JRZMrJWFZZnZtJgvvjwbpbCJDd
	rFMPRhZtZFnWrRtQGmPPDcLfmGLTfz
	VdWnVdjhhdFjVWbndMlNLQspVMHCNVlClV
	RSrJBRRJwJSBQpMBHLLDCL
	TqwtRRRJzJTSqJSzSrtmqgWWhcncvPgnWbPQnbnWmb
	VnDFpPpFssVSpFDVHbRbscCvgbMTvTCR
	JfzqdQBfhBdddfBBGDLdGQvbrqMMcCRRMTgbqgMrbbqc
	QDfzJNWBJLQBhmdGDzDGhQGGlFZwPtWjtFFppllSVpZZFnjj
	qrLLNpJbJnRLNnpvQtRVhhRFCdlFFlFd
	mmjzjvGjwPwmTsSTSQjDVlVWQjlCDthCCC
	cSSmcTTPcSswScfSHmTSTzJqqNrnpBpqBbJLvZMrqfrL
	NSvRZRfFvfHSZQcNJBLbzDLnrDFnhtFLFnrh
	wmTGpmGCwsMplMsHllPlMnDLjznrgrzDjgnntznr
	dsCVGGGwmpTGPplmCmPppVmHSSRJNfJvBNZQfWdBRJZRBZcR
	TwQwqDPQtwNwzNDTZcnZbJvMnMMbFqZM
	SzGSjrjLWrjHHspWVhvVVnFJbccVZcRJbllb
	pHppszGSprhhWHLCLrsjdTtDDPfwdfwtdNfDgNCN
	ftcvBtBFtmBlmvPFmmcczCChrgSCzzCSnCSSnGHf
	sJddbdTDbDHdnJRggrGzGzrG
	dppDVDZMMMsTTVsDTsTDpwVctNcvBZQPcPctqtQcHmvlvQ
	jzbdzztbDqNqwvLvRmQZjvRH
	FSJbFFWgJnZFLRZmHmRQ
	TgVJTVSJGJcJlllgTMdqpdNsrztNNsNbMDDp
	CCCVWbwVnlRbTcqSShqGhhGcnF
	PgDBfDpMNlfgpPfNZZtcJgcqqhmmjqSmjFmhmS
	tpfpsPrlpsPDDMDfBZrwLrVWLLLWRCdHLTwbVR
	pjvfDGjSMpvDmDpDpSDnJmfqbPVsCMFsPqFVPqCrwrbMFV
	NQlHtHNhZHgZZNBHhQgzPmCwbqqVFlsrPFrFCP
	hgHtQdQchcHctHgcgNgBQdWNpmvTWvpGmLJDLGjTpLGnnjfv
	QhgLLLmtlRqDtRGP
	HLbnCZFWVHLZnFCJJRFrGJzDGDGJDD
	WZHfndfMfCZbMnTVTfZhSNQQpdwSdLwhNcmdSN
	sPwrPMgLFPFFsLZtmcclSSZDtcZs
	qVzqdNdCnnNVVNCGmbncDBlmBlBBnRlZ
	VTdCGVvVfffrjpfMQPwm
	BPDldDTDPZcggjcccTdNMbbMNSQNqqjtzMbrRb
	LvmWsfvssLGnQbQMRQqrSRnz
	WpvsVmmpmmfpfJGrHfVCHVvmcDgpDlZphgFgdhclhdgdBlgF
	VGwHbNzMMrzHbbHChhqgCqPNghgCqW
	ZJVBvBvZWqvRvggP
	JBJlBlBZcsBfcJVrHnLwQQGzLQMc
	gBWfBPPPfhvVWFfSVfVdjjbvTvwwQppHcHcctTcQTHcZ
	DnNnMJMqMJzqchbZtTQQrb
	llRmNLDLDGlCsWSFCffWdshd
	LpNMZZpqqpfTTwNqLZwGsZqZbdHRHbHGddnCBHRcmzGmmCdG
	JFRtRlVStjPlhtjbBzBncmVWdzWBnb
	rPhhSlrvQlFFFPgtJlJtFlhlDNTwRMfZTZfDZNrspZLMMsrq
	zBLjLFBjLjmHWlzNZlzVCC
	dcJrdfddbllJbdMTwDNMZWNVwVDwHT
	gRcgJbcbqfgbftdjlqLhFFLPPhGBjm
	WfBgBRzQGNNQqmmqZN
	nFjCjCpLbtpPJtCDDnCDJpzncrSVbmdVqbhhdqNbSSmrdVSq
	CLPJpDLlLlFDpFjjsGRsBGRfWwsHHglz
	lSlSlpCRSsWTRLTlWRvlmMrBPjBPjpqrrmqPJMPZ
	DDzbhVhQhDGzhQnGGfnHHQGBPZjMqJjBJMBVJmqMdrqqdT
	NNGQbFwnHzNzwbQwFnwbfsLCLtsvLsWggFslsTggSc
	nvzPvCnlvtwCrZWmWwvvZCQfbbfQfGbqSJJGmqGSFSbJ
	LhTBWdsMNNRgNcgDWsDNcVSfQqJGFSFJqSSddQGSJF
	HNgchHcWDRNhTNMWwtPrtZZjnHzrnvCz
	djhnzRghMMVCBfhh
	qjQTrTPQJCDDqBDJ
	LQvGrLjTHLjNNPPTpQgtztSmmbFgmgLbFnmL
	FRDNFBBRRVFFmbLZHPZBZvvH
	QnhgMllglJTdGgJnhLQQJpZpvwZHpwsPTwpbsZHmsH
	lnhnQGrMgthMlntlGfQhgWWcRSDcVCrLWzRSrRFDRN
	PqrrrRnPBbrVhVqFrFVRPVhZLvNSNvLZcQvtJfRvNScJNJ
	dDzWwwCTmmdwdddpDLWQZMSSMfSJtcWJfQSQZN
	CCwmTdjsClVjFjnLBl
	srjCvjPmQVlPjFPmQmPrdHHZhvHZDqHhDDwHHqfB
	pLcnJQNQMZpqZDDZ
	WNRbtNJgRPjjQVmz
	NJJRmjmJbbJfqSVMNHFCSFzLLlrLLrFHTz
	QvnsQGvBwWwQvgRHlGGDFPFCGlrR
	QhvwBvBctBccZWZNRNmVfjpmjJjb
	RMmGGMLRRCFmRPPfGFpGPFPJWZQWctrtlQvZvltfrQWcWWBq
	gggwjjbjwwbZtwZBBcmQQv
	SdNbDDVSgPMFmPzdMm
	nZhnNZDnZPmZPWbppPpMlvRlzvrtMmRtqRzRfq
	HcFwsCQLVQwFwLtLbvtzrlrLtt
	GsgCFCgCQHHCVHsFQHcFdDPDbJDZTpZDbWJPNWWZDd
	BBrBrGlGpgGjsNhlBlpBwpfSwZJdQwfcZwvSQnnn
	LvWvHLmmVJQQHfQH
	RPLRMvqFTbRTjGBhjNFsslls
	cNZZZmZDcDDJmhzzrrlHtSbvgjSvgfPSWvPfjShv
	VBwnndnVCqbqpRRpnspnqRWtGgWSSgvFBSGGSWgtGGSP
	LqCMnTLVRwCRCpRLpbHDNzMMNcmmHNHQJQ
	MMqDtnVnBlHtZvtB
	WLWrWgdWwdrLCTFCwLlbbsJsJQsbQlQzlvrB
	jFSvTdjfnfRmVcRR
	ZLGqnvnqLzvbGRMfcRpwMpdV
	fgfNNfgHHjVmRcVdgM
	HsWDCDfCQCZBBZnvWtLq
	bTZjqflqZhcrlczGzppGNgjmFNnp
	PmmRSWWDMBQVNpWFznGF
	SStRBDSCCSSSwPBwBDBwPmZhZlfZhqHTsTfltHHZfsHH
	GbNbsSptQGqsdJCzsddcgzzv
	DHRRnmWWmZnmRhllnHnnnMLvvLgcTVvjVhCTvgzcJgLj
	RnWMlDZRlnHlmHWBFwGQqNGGPNQzPGqFwz
	vSGvHpJnBLbGHBNCgfDzzChDgbCfzT
	wFRslqmqTRgggQghPmQf
	qjRFMjWqNNMMGHTL
	fWGcQGGSRFQZhttZJfSSJflDDrwdClljVrNDdrdCFBCr
	MTgvLLPPnHzMbDwdlNbMBwMM
	mnTvnnPTcNmmJJWN
	qqbbQQnbWrqGgnWqvZpVzMCZjCgfjZCSVM
	ldcmDPDhmlFBHPDddLBVFDHLppZpjSCjjNfwNMwCpSMwhCMp
	FtDdsHPcHmdHVPLtHsdtBHQnsbvnTRRTRsRRqbqvqWnJ
	hhtBtPrgbbhhgjZjjCCHHNpNDHpffHWCvr
	LGFLVwswsJMSgFwMMpddSvpHCCdDdvCpvm
	sGsFsQLsVsLFnnFTJQthjcjQqhRcBZZtRg`


	//phase 1
			// var rucksA []Rucksack
			// var rucksB []RucksackDups
			// //one item per group
			// var rucksC []RucksackDupsValue
			// var sum int

			// //turn rows into rucksacks
			// _, rucks1 := computeInput(input,rucksA)

			// //turn rucksacks into rucksacks with duplicates
			// for i := range rucks1{
			// 	rucksB = append(rucksB, findDups(rucks1[i]))
			// }
			// //turn rucksack dups into ruck sacks with values
			// for j := range rucksB{
			// 		rucksC = append(rucksC, determineValueofDup(rucksB[j]))
			
			// }
			// //turn rucksacks with duplicates into rucksacks with values
			// for k := range rucksC{
			// 	sum = sum + rucksC[k].value
			// }

			// fmt.Println("sum =", sum)

			// for m := range rucksC{
			// 	fmt.Println("--------------------------------------------------")
			// 	fmt.Println("m = ", m )
			// 	fmt.Println("Dup found in each row = ", string(rucksC[m].Dup))
			// 	fmt.Println("Value  of each dup found in each row = ", rucksC[m].value)
			// 	fmt.Println("--------------------------------------------------")
			// }

	//phase 2:

			_, total := determineValueInGroup(input, 0)
			fmt.Println(total)
}
func computeInput (str string, Rucks []Rucksack) (string, []Rucksack){
rucks := Rucks
 	var ruck Rucksack
	str = strings.TrimSpace(str)
	for i:= range str{
		
		if(str[i]) ==10{
			temp := str[:i]
			length := len(temp)
			half := length/2
			Com1 := temp[:half]
			Com2 := temp[half:]
			ruck.Com1=Com1
			ruck.Com2 = Com2
			rucks = append(rucks, ruck)
			return computeInput(strings.TrimSpace(str[i+1:]), rucks)
		}
	}
	temp := str
	length := len(temp)
	half := length/2
	Com1 := temp[:half]
	Com2 := temp[half+1:length]
	ruck.Com1=Com1
	ruck.Com2 = Com2
	// fmt.Println("Com 1 = ", Com1)
	// fmt.Println("Com2 = ", Com2)
	rucks = append(rucks, ruck)
return "", rucks
}

func findDups (ruck Rucksack) RucksackDups{

	first := ruck.Com1
	second := ruck.Com2
	// fmt.Println("first com = ", first)
	// fmt.Println("second com = ", second)
	var ruckDups RucksackDups

	for i := range first{
		for j := range second{
			if(first[i]==second[j]){
				ruckDups.Com1 = first
				ruckDups.Com2 = second
				ruckDups.Dup = rune(first[i])
				return ruckDups
			}
		}
	}
	return ruckDups
}

func determineValueofDup (ruck RucksackDups) RucksackDupsValue{
	var rv RucksackDupsValue
	var value int
	bt := byte(ruck.Dup)
	if int(bt) > 96{
		value = int(bt)-96
	}
	if int(bt) <= 96{
		value = int(bt) - 38
	}

	rv.Com1 = ruck.Com1
	rv.Com2 = ruck.Com2
	rv.value = value
	rv.Dup = ruck.Dup
	return rv
}


func determineValueInGroup(str string, total int) (string, int){
	rowCount :=0
	var indexList []int
	for i:= range str{


		if(str[i]) ==10{
			rowCount = rowCount +1
			indexList = append(indexList, i)
			fmt.Println("Adding an index. Count is = ", len(indexList))
		}

		if(rowCount ==3){
			temp := str[:i]
			first := strings.TrimSpace(temp[:indexList[0]])
			second := strings.TrimSpace(temp[indexList[0]:indexList[1]])
			third := strings.TrimSpace(temp[indexList[1]:indexList[2]])
			fmt.Println("-----We've hit a group--------")
			fmt.Println("first = ", first)
			fmt.Println("second = ", second)
			fmt.Println("third = ", third)
			total = dupValueThirds(first,second,third) + total
			fmt.Println("total = ", total)
			
			// fmt.Println("Com 1 = ", Com1)
			// fmt.Println("Com2 = ", Com2)
			return determineValueInGroup(str[indexList[2]+1:], total)
		}
	}
	first := strings.TrimSpace(str[:indexList[0]])
	second := strings.TrimSpace(str[indexList[0]:indexList[1]])
	third := strings.TrimSpace(str[indexList[1]:])
	fmt.Println("-----We've hit a group--------")
	fmt.Println("first = ", first)
	fmt.Println("second = ", second)
	fmt.Println("third = ", third)
	total = dupValueThirds(first,second,third) + total
	fmt.Println("total = ", total)
	fmt.Println("------------------------------")
	// fmt.Println("Com 1 = ", Com1)
	// fmt.Println("Com2 = ", Com2)
return "", 0
}

func dupValueThirds(first string, second string, third string) int{
	for i := range first{
		for j := range second{
			for l := range third{
				if first[i] == second[j] && second[j] == third[l]{
					fmt.Println("dup is =", string(first[i]))
					fmt.Println("------------------------------")
					bt := byte(first[i])
					if int(bt) > 96{
						return  int(bt)-96
					}
					if int(bt) <= 96{
						return int(bt) - 38
					}
				}
			}
		}
	}
	return -100000
}