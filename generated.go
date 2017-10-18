// Copyright 2017 Dale Farnsworth. All rights reserved.

// Dale Farnsworth
// 1007 W Mendoza Ave
// Mesa, AZ  85210
// USA
//
// dale@farnsworth.org

// This file is part of Codeplug.
//
// Codeplug is free software: you can redistribute it and/or modify
// it under the terms of version 3 of the GNU Lesser General Public
// License as published by the Free Software Foundation.
//
// Codeplug is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Codeplug.  If not, see <http://www.gnu.org/licenses/>.

// Package codeplug implements access to MD380-style codeplug files.
// It can read/update/write both .rdt files and .bin files.
package codeplug

//go:generate genCodeplugInfo codeplugs.json

// Record types
const (
	RtChannelInformation    RecordType = "ChannelInformation"
	RtDigitalContacts       RecordType = "DigitalContacts"
	RtGeneralSettings       RecordType = "GeneralSettings"
	RtGroupList             RecordType = "GroupList"
	RtRdtHeader             RecordType = "RdtHeader"
	RtScanList              RecordType = "ScanList"
	RtTextMessage           RecordType = "TextMessage"
	RtZoneInformation_md380 RecordType = "ZoneInformation_md380"
	RtZoneInformation_md40  RecordType = "ZoneInformation_md40"
)

// Field types
const (
	FtCiAdmitCriteria           FieldType = "CiAdmitCriteria"
	FtCiAllowTalkaround         FieldType = "CiAllowTalkaround"
	FtCiAutoscan                FieldType = "CiAutoscan"
	FtCiBandwidth               FieldType = "CiBandwidth"
	FtCiChannelMode             FieldType = "CiChannelMode"
	FtCiChannelName             FieldType = "CiChannelName"
	FtCiColorCode               FieldType = "CiColorCode"
	FtCiCompressedUdpDataHeader FieldType = "CiCompressedUdpDataHeader"
	FtCiContactName             FieldType = "CiContactName"
	FtCiCtcssDecode             FieldType = "CiCtcssDecode"
	FtCiCtcssEncode             FieldType = "CiCtcssEncode"
	FtCiDataCallConfirmed       FieldType = "CiDataCallConfirmed"
	FtCiDecode1                 FieldType = "CiDecode1"
	FtCiDecode2                 FieldType = "CiDecode2"
	FtCiDecode3                 FieldType = "CiDecode3"
	FtCiDecode4                 FieldType = "CiDecode4"
	FtCiDecode5                 FieldType = "CiDecode5"
	FtCiDecode6                 FieldType = "CiDecode6"
	FtCiDecode7                 FieldType = "CiDecode7"
	FtCiDecode8                 FieldType = "CiDecode8"
	FtCiDisplayPTTID            FieldType = "CiDisplayPTTID"
	FtCiEmergencyAlarmAck       FieldType = "CiEmergencyAlarmAck"
	FtCiGroupList               FieldType = "CiGroupList"
	FtCiLoneWorker              FieldType = "CiLoneWorker"
	FtCiPower                   FieldType = "CiPower"
	FtCiPrivacy                 FieldType = "CiPrivacy"
	FtCiPrivacyNumber           FieldType = "CiPrivacyNumber"
	FtCiPrivateCallConfirmed    FieldType = "CiPrivateCallConfirmed"
	FtCiQtReverse               FieldType = "CiQtReverse"
	FtCiRepeaterSlot            FieldType = "CiRepeaterSlot"
	FtCiReverseBurst            FieldType = "CiReverseBurst"
	FtCiRxFrequency             FieldType = "CiRxFrequency"
	FtCiRxOnly                  FieldType = "CiRxOnly"
	FtCiRxRefFrequency          FieldType = "CiRxRefFrequency"
	FtCiRxSignallingSystem      FieldType = "CiRxSignallingSystem"
	FtCiScanList                FieldType = "CiScanList"
	FtCiSquelch                 FieldType = "CiSquelch"
	FtCiTot                     FieldType = "CiTot"
	FtCiTotRekeyDelay           FieldType = "CiTotRekeyDelay"
	FtCiTxFrequency             FieldType = "CiTxFrequency"
	FtCiTxRefFrequency          FieldType = "CiTxRefFrequency"
	FtCiTxSignallingSystem      FieldType = "CiTxSignallingSystem"
	FtCiVox                     FieldType = "CiVox"
	FtDcCallID                  FieldType = "DcCallID"
	FtDcCallReceiveTone         FieldType = "DcCallReceiveTone"
	FtDcCallType                FieldType = "DcCallType"
	FtDcContactName             FieldType = "DcContactName"
	FtGlContactMember           FieldType = "GlContactMember"
	FtGlName                    FieldType = "GlName"
	FtGsBacklightColor          FieldType = "GsBacklightColor"
	FtGsBacklightTime           FieldType = "GsBacklightTime"
	FtGsCallAlertToneDuration   FieldType = "GsCallAlertToneDuration"
	FtGsChFreeIndicationTone    FieldType = "GsChFreeIndicationTone"
	FtGsDisableAllLeds          FieldType = "GsDisableAllLeds"
	FtGsDisableAllTones         FieldType = "GsDisableAllTones"
	FtGsFreqChannelMode         FieldType = "GsFreqChannelMode"
	FtGsGroupCallHangTime       FieldType = "GsGroupCallHangTime"
	FtGsIntroScreen             FieldType = "GsIntroScreen"
	FtGsIntroScreenLine1        FieldType = "GsIntroScreenLine1"
	FtGsIntroScreenLine2        FieldType = "GsIntroScreenLine2"
	FtGsLockUnlock              FieldType = "GsLockUnlock"
	FtGsLoneWorkerReminderTime  FieldType = "GsLoneWorkerReminderTime"
	FtGsLoneWorkerResponseTime  FieldType = "GsLoneWorkerResponseTime"
	FtGsMode                    FieldType = "GsMode"
	FtGsModeSelect              FieldType = "GsModeSelect"
	FtGsMonitorType             FieldType = "GsMonitorType"
	FtGsPcProgPw                FieldType = "GsPcProgPw"
	FtGsPowerOnPassword         FieldType = "GsPowerOnPassword"
	FtGsPrivateCallHangTime     FieldType = "GsPrivateCallHangTime"
	FtGsPwAndLockEnable         FieldType = "GsPwAndLockEnable"
	FtGsRadioID                 FieldType = "GsRadioID"
	FtGsRadioName               FieldType = "GsRadioName"
	FtGsRadioProgPw             FieldType = "GsRadioProgPw"
	FtGsRxLowBatteryInterval    FieldType = "GsRxLowBatteryInterval"
	FtGsSaveModeReceive         FieldType = "GsSaveModeReceive"
	FtGsSavePreamble            FieldType = "GsSavePreamble"
	FtGsScanAnalogHangTime      FieldType = "GsScanAnalogHangTime"
	FtGsScanDigitalHangTime     FieldType = "GsScanDigitalHangTime"
	FtGsSetKeypadLockTime       FieldType = "GsSetKeypadLockTime"
	FtGsTalkPermitTone          FieldType = "GsTalkPermitTone"
	FtGsTxPreambleDuration      FieldType = "GsTxPreambleDuration"
	FtGsVoxSensitivity          FieldType = "GsVoxSensitivity"
	FtRhHighFrequency           FieldType = "RhHighFrequency"
	FtRhLabel                   FieldType = "RhLabel"
	FtRhLowFrequency            FieldType = "RhLowFrequency"
	FtSlChannelMember           FieldType = "SlChannelMember"
	FtSlName                    FieldType = "SlName"
	FtSlPriorityChannel1        FieldType = "SlPriorityChannel1"
	FtSlPriorityChannel2        FieldType = "SlPriorityChannel2"
	FtSlPrioritySampleTime      FieldType = "SlPrioritySampleTime"
	FtSlSignallingHoldTime      FieldType = "SlSignallingHoldTime"
	FtSlTxDesignatedChannel     FieldType = "SlTxDesignatedChannel"
	FtTmTextMessage             FieldType = "TmTextMessage"
	FtZiChannelMember_md380     FieldType = "ZiChannelMember_md380"
	FtZiChannelMember_md40      FieldType = "ZiChannelMember_md40"
	FtZiName                    FieldType = "ZiName"
)

// The value types a field may contain
const (
	VtAscii           ValueType = "ascii"
	VtCallID          ValueType = "callID"
	VtCtcssDcs        ValueType = "ctcssDcs"
	VtFrequency       ValueType = "frequency"
	VtIStrings        ValueType = "iStrings"
	VtIndexedStrings  ValueType = "indexedStrings"
	VtIntroLine       ValueType = "introLine"
	VtListIndex       ValueType = "listIndex"
	VtMemberListIndex ValueType = "memberListIndex"
	VtName            ValueType = "name"
	VtOffOn           ValueType = "offOn"
	VtOnOff           ValueType = "onOff"
	VtPcPassword      ValueType = "pcPassword"
	VtPrivacyNumber   ValueType = "privacyNumber"
	VtRadioName       ValueType = "radioName"
	VtRadioPassword   ValueType = "radioPassword"
	VtRhFrequency     ValueType = "rhFrequency"
	VtSpan            ValueType = "span"
	VtTextMessage     ValueType = "textMessage"
)

// newValue returns a new value of the given ValueType
func newValue(vt ValueType) value {
	switch vt {
	case VtAscii:
		return new(ascii)
	case VtCallID:
		return new(callID)
	case VtCtcssDcs:
		return new(ctcssDcs)
	case VtFrequency:
		return new(frequency)
	case VtIStrings:
		return new(iStrings)
	case VtIndexedStrings:
		return new(indexedStrings)
	case VtIntroLine:
		return new(introLine)
	case VtListIndex:
		return new(listIndex)
	case VtMemberListIndex:
		return new(memberListIndex)
	case VtName:
		return new(name)
	case VtOffOn:
		return new(offOn)
	case VtOnOff:
		return new(onOff)
	case VtPcPassword:
		return new(pcPassword)
	case VtPrivacyNumber:
		return new(privacyNumber)
	case VtRadioName:
		return new(radioName)
	case VtRadioPassword:
		return new(radioPassword)
	case VtRhFrequency:
		return new(rhFrequency)
	case VtSpan:
		return new(span)
	case VtTextMessage:
		return new(textMessage)
	}

	return nil
}

var codeplugInfos = []*CodeplugInfo{
	&cpMd380,
	&cpMd40,
}

var cpMd380 = CodeplugInfo{
	Name: "MD380/MD390/DR780",
	Type: "md380",
	Labels: []string{
		"MD380",
		"MD390",
		"DR780",
	},
	RdtSize:   262709,
	BinSize:   262144,
	BinOffset: 549,
	RecordInfos: []*recordInfo{
		&riRdtHeader,
		&riGeneralSettings,
		&riTextMessage,
		&riDigitalContacts,
		&riGroupList,
		&riZoneInformation_md380,
		&riScanList,
		&riChannelInformation,
	},
}

var cpMd40 = CodeplugInfo{
	Name: "DJ-MD40",
	Type: "md40",
	Labels: []string{
		"DJ-MD40",
	},
	RdtSize:   262709,
	BinSize:   262144,
	BinOffset: 549,
	RecordInfos: []*recordInfo{
		&riRdtHeader,
		&riGeneralSettings,
		&riTextMessage,
		&riDigitalContacts,
		&riGroupList,
		&riZoneInformation_md40,
		&riScanList,
		&riChannelInformation,
	},
}

var riChannelInformation = recordInfo{
	rType:    RtChannelInformation,
	typeName: "Channel Information",
	max:      1000,
	offset:   127013,
	size:     64,
	delDescs: []delDesc{
		delDesc{
			offset: 16,
			size:   1,
			value:  255,
		},
	},
	fieldInfos: []*fieldInfo{
		&fiCiLoneWorker,
		&fiCiSquelch,
		&fiCiAutoscan,
		&fiCiBandwidth,
		&fiCiChannelMode,
		&fiCiColorCode,
		&fiCiRepeaterSlot,
		&fiCiRxOnly,
		&fiCiAllowTalkaround,
		&fiCiDataCallConfirmed,
		&fiCiPrivateCallConfirmed,
		&fiCiPrivacy,
		&fiCiPrivacyNumber,
		&fiCiDisplayPTTID,
		&fiCiCompressedUdpDataHeader,
		&fiCiEmergencyAlarmAck,
		&fiCiRxRefFrequency,
		&fiCiAdmitCriteria,
		&fiCiPower,
		&fiCiVox,
		&fiCiQtReverse,
		&fiCiReverseBurst,
		&fiCiTxRefFrequency,
		&fiCiContactName,
		&fiCiTot,
		&fiCiTotRekeyDelay,
		&fiCiScanList,
		&fiCiGroupList,
		&fiCiDecode1,
		&fiCiDecode2,
		&fiCiDecode3,
		&fiCiDecode4,
		&fiCiDecode5,
		&fiCiDecode6,
		&fiCiDecode7,
		&fiCiDecode8,
		&fiCiRxFrequency,
		&fiCiTxFrequency,
		&fiCiCtcssDecode,
		&fiCiCtcssEncode,
		&fiCiRxSignallingSystem,
		&fiCiTxSignallingSystem,
		&fiCiChannelName,
	},
}

var riDigitalContacts = recordInfo{
	rType:    RtDigitalContacts,
	typeName: "Digital Contacts",
	max:      1000,
	offset:   24997,
	size:     36,
	delDescs: []delDesc{
		delDesc{
			offset: 0,
			size:   3,
			value:  255,
		},
		delDesc{
			offset: 4,
			size:   2,
			value:  0,
		},
		delDesc{
			offset: 4,
			size:   16,
			value:  0,
		},
	},
	fieldInfos: []*fieldInfo{
		&fiDcCallID,
		&fiDcCallReceiveTone,
		&fiDcCallType,
		&fiDcContactName,
	},
}

var riGeneralSettings = recordInfo{
	rType:    RtGeneralSettings,
	typeName: "General Settings",
	max:      1,
	offset:   8805,
	size:     144,
	fieldInfos: []*fieldInfo{
		&fiGsIntroScreenLine1,
		&fiGsIntroScreenLine2,
		&fiGsMonitorType,
		&fiGsDisableAllLeds,
		&fiGsTalkPermitTone,
		&fiGsPwAndLockEnable,
		&fiGsChFreeIndicationTone,
		&fiGsDisableAllTones,
		&fiGsSaveModeReceive,
		&fiGsSavePreamble,
		&fiGsIntroScreen,
		&fiGsLockUnlock,
		&fiGsFreqChannelMode,
		&fiGsModeSelect,
		&fiGsBacklightColor,
		&fiGsRadioID,
		&fiGsTxPreambleDuration,
		&fiGsGroupCallHangTime,
		&fiGsPrivateCallHangTime,
		&fiGsVoxSensitivity,
		&fiGsRxLowBatteryInterval,
		&fiGsCallAlertToneDuration,
		&fiGsLoneWorkerResponseTime,
		&fiGsLoneWorkerReminderTime,
		&fiGsScanDigitalHangTime,
		&fiGsScanAnalogHangTime,
		&fiGsBacklightTime,
		&fiGsSetKeypadLockTime,
		&fiGsMode,
		&fiGsPowerOnPassword,
		&fiGsRadioProgPw,
		&fiGsPcProgPw,
		&fiGsRadioName,
	},
}

var riGroupList = recordInfo{
	rType:    RtGroupList,
	typeName: "Digital Rx Group List",
	max:      250,
	offset:   60997,
	size:     96,
	delDescs: []delDesc{
		delDesc{
			offset: 0,
			size:   1,
			value:  0,
		},
	},
	fieldInfos: []*fieldInfo{
		&fiGlName,
		&fiGlContactMember,
	},
}

var riRdtHeader = recordInfo{
	rType:    RtRdtHeader,
	typeName: "Rdt Header",
	max:      1,
	offset:   0,
	size:     549,
	fieldInfos: []*fieldInfo{
		&fiRhLabel,
		&fiRhLowFrequency,
		&fiRhHighFrequency,
	},
}

var riScanList = recordInfo{
	rType:    RtScanList,
	typeName: "Scan List",
	max:      250,
	offset:   100997,
	size:     104,
	delDescs: []delDesc{
		delDesc{
			offset: 0,
			size:   1,
			value:  0,
		},
	},
	fieldInfos: []*fieldInfo{
		&fiSlName,
		&fiSlPriorityChannel1,
		&fiSlPriorityChannel2,
		&fiSlTxDesignatedChannel,
		&fiSlSignallingHoldTime,
		&fiSlPrioritySampleTime,
		&fiSlChannelMember,
	},
}

var riTextMessage = recordInfo{
	rType:    RtTextMessage,
	typeName: "Text Message",
	max:      50,
	offset:   9125,
	size:     288,
	delDescs: []delDesc{
		delDesc{
			offset: 0,
			size:   8,
			value:  0,
		},
	},
	fieldInfos: []*fieldInfo{
		&fiTmTextMessage,
	},
}

var riZoneInformation_md380 = recordInfo{
	rType:    RtZoneInformation_md380,
	typeName: "Zone Information",
	max:      250,
	offset:   84997,
	size:     64,
	delDescs: []delDesc{
		delDesc{
			offset: 0,
			size:   1,
			value:  0,
		},
	},
	fieldInfos: []*fieldInfo{
		&fiZiName,
		&fiZiChannelMember_md380,
	},
}

var riZoneInformation_md40 = recordInfo{
	rType:    RtZoneInformation_md40,
	typeName: "Zone Information",
	max:      250,
	offset:   84997,
	size:     64,
	delDescs: []delDesc{
		delDesc{
			offset: 0,
			size:   1,
			value:  0,
		},
	},
	fieldInfos: []*fieldInfo{
		&fiZiName,
		&fiZiChannelMember_md40,
	},
}

var fiCiAdmitCriteria = fieldInfo{
	fType:     FtCiAdmitCriteria,
	typeName:  "Admit Criteria",
	max:       1,
	bitOffset: 32,
	bitSize:   2,
	valueType: VtIStrings,
	strings: &[]string{
		"Always",
		"Channel free",
		"CTCSS/DCS",
		"Color code",
	},
}

var fiCiAllowTalkaround = fieldInfo{
	fType:     FtCiAllowTalkaround,
	typeName:  "Allow Talkaround",
	max:       1,
	bitOffset: 15,
	bitSize:   1,
	valueType: VtOffOn,
}

var fiCiAutoscan = fieldInfo{
	fType:     FtCiAutoscan,
	typeName:  "Autoscan",
	max:       1,
	bitOffset: 3,
	bitSize:   1,
	valueType: VtOffOn,
}

var fiCiBandwidth = fieldInfo{
	fType:     FtCiBandwidth,
	typeName:  "Bandwidth",
	max:       1,
	bitOffset: 4,
	bitSize:   1,
	valueType: VtIStrings,
	strings: &[]string{
		"12.5",
		"25",
	},
}

var fiCiChannelMode = fieldInfo{
	fType:     FtCiChannelMode,
	typeName:  "Channel Mode",
	max:       1,
	bitOffset: 6,
	bitSize:   2,
	valueType: VtIStrings,
	strings: &[]string{
		"",
		"Analog",
		"Digital",
	},
	enablingValue: "Digital",
}

var fiCiChannelName = fieldInfo{
	fType:     FtCiChannelName,
	typeName:  "Channel Name",
	max:       1,
	bitOffset: 256,
	bitSize:   256,
	valueType: VtName,
}

var fiCiColorCode = fieldInfo{
	fType:     FtCiColorCode,
	typeName:  "Color Code",
	max:       1,
	bitOffset: 8,
	bitSize:   4,
	valueType: VtSpan,
	span: &Span{
		min: 0,
		max: 15,
	},
	enabler: FtCiChannelMode,
}

var fiCiCompressedUdpDataHeader = fieldInfo{
	fType:     FtCiCompressedUdpDataHeader,
	typeName:  "Compressed UDP Data Header",
	max:       1,
	bitOffset: 25,
	bitSize:   1,
	valueType: VtOffOn,
	enabler:   FtCiChannelMode,
}

var fiCiContactName = fieldInfo{
	fType:        FtCiContactName,
	typeName:     "Contact Name",
	max:          1,
	bitOffset:    48,
	bitSize:      16,
	valueType:    VtListIndex,
	defaultValue: "None",
	indexedStrings: &[]IndexedString{
		IndexedString{0, "None"},
	},
	listRecordType: RtDigitalContacts,
	enabler:        FtCiChannelMode,
}

var fiCiCtcssDecode = fieldInfo{
	fType:        FtCiCtcssDecode,
	typeName:     "CTCSS/DCS Decode",
	max:          1,
	bitOffset:    192,
	bitSize:      16,
	valueType:    VtCtcssDcs,
	defaultValue: "None",
	disabler:     FtCiChannelMode,
}

var fiCiCtcssEncode = fieldInfo{
	fType:         FtCiCtcssEncode,
	typeName:      "CTCSS/DCS Encode",
	max:           1,
	bitOffset:     208,
	bitSize:       16,
	valueType:     VtCtcssDcs,
	defaultValue:  "None",
	enablingValue: "None",
	disabler:      FtCiChannelMode,
}

var fiCiDataCallConfirmed = fieldInfo{
	fType:     FtCiDataCallConfirmed,
	typeName:  "Data Call Confirmed",
	max:       1,
	bitOffset: 16,
	bitSize:   1,
	valueType: VtOffOn,
	enabler:   FtCiChannelMode,
}

var fiCiDecode1 = fieldInfo{
	fType:     FtCiDecode1,
	typeName:  "Decode 1",
	max:       1,
	bitOffset: 112,
	bitSize:   1,
	valueType: VtOffOn,
	disabler:  FtCiRxSignallingSystem,
}

var fiCiDecode2 = fieldInfo{
	fType:     FtCiDecode2,
	typeName:  "Decode 2",
	max:       1,
	bitOffset: 113,
	bitSize:   1,
	valueType: VtOffOn,
	disabler:  FtCiRxSignallingSystem,
}

var fiCiDecode3 = fieldInfo{
	fType:     FtCiDecode3,
	typeName:  "Decode 3",
	max:       1,
	bitOffset: 114,
	bitSize:   1,
	valueType: VtOffOn,
	disabler:  FtCiRxSignallingSystem,
}

var fiCiDecode4 = fieldInfo{
	fType:     FtCiDecode4,
	typeName:  "Decode 4",
	max:       1,
	bitOffset: 115,
	bitSize:   1,
	valueType: VtOffOn,
	disabler:  FtCiRxSignallingSystem,
}

var fiCiDecode5 = fieldInfo{
	fType:     FtCiDecode5,
	typeName:  "Decode 5",
	max:       1,
	bitOffset: 116,
	bitSize:   1,
	valueType: VtOffOn,
	disabler:  FtCiRxSignallingSystem,
}

var fiCiDecode6 = fieldInfo{
	fType:     FtCiDecode6,
	typeName:  "Decode 6",
	max:       1,
	bitOffset: 117,
	bitSize:   1,
	valueType: VtOffOn,
	disabler:  FtCiRxSignallingSystem,
}

var fiCiDecode7 = fieldInfo{
	fType:     FtCiDecode7,
	typeName:  "Decode 7",
	max:       1,
	bitOffset: 118,
	bitSize:   1,
	valueType: VtOffOn,
	disabler:  FtCiRxSignallingSystem,
}

var fiCiDecode8 = fieldInfo{
	fType:     FtCiDecode8,
	typeName:  "Decode 8",
	max:       1,
	bitOffset: 119,
	bitSize:   1,
	valueType: VtOffOn,
	disabler:  FtCiRxSignallingSystem,
}

var fiCiDisplayPTTID = fieldInfo{
	fType:     FtCiDisplayPTTID,
	typeName:  "Display PTT ID",
	max:       1,
	bitOffset: 24,
	bitSize:   1,
	valueType: VtOnOff,
	disabler:  FtCiChannelMode,
}

var fiCiEmergencyAlarmAck = fieldInfo{
	fType:     FtCiEmergencyAlarmAck,
	typeName:  "Emergency Alarm Ack",
	max:       1,
	bitOffset: 28,
	bitSize:   1,
	valueType: VtOffOn,
	enabler:   FtCiChannelMode,
}

var fiCiGroupList = fieldInfo{
	fType:        FtCiGroupList,
	typeName:     "Group List",
	max:          1,
	bitOffset:    96,
	bitSize:      8,
	valueType:    VtListIndex,
	defaultValue: "None",
	indexedStrings: &[]IndexedString{
		IndexedString{0, "None"},
	},
	listRecordType: RtGroupList,
	enabler:        FtCiChannelMode,
}

var fiCiLoneWorker = fieldInfo{
	fType:     FtCiLoneWorker,
	typeName:  "Lone Worker",
	max:       1,
	bitOffset: 0,
	bitSize:   1,
	valueType: VtOffOn,
}

var fiCiPower = fieldInfo{
	fType:     FtCiPower,
	typeName:  "Power",
	max:       1,
	bitOffset: 34,
	bitSize:   1,
	valueType: VtIStrings,
	strings: &[]string{
		"Low",
		"High",
	},
}

var fiCiPrivacy = fieldInfo{
	fType:        FtCiPrivacy,
	typeName:     "Privacy",
	max:          1,
	bitOffset:    18,
	bitSize:      2,
	valueType:    VtIStrings,
	defaultValue: "None",
	strings: &[]string{
		"None",
		"Basic",
		"Enhanced",
	},
	enablingValue: "None",
	enabler:       FtCiChannelMode,
}

var fiCiPrivacyNumber = fieldInfo{
	fType:        FtCiPrivacyNumber,
	typeName:     "Privacy Number",
	max:          1,
	bitOffset:    20,
	bitSize:      4,
	valueType:    VtPrivacyNumber,
	defaultValue: "0",
	span: &Span{
		min: 0,
		max: 15,
	},
	disabler: FtCiPrivacy,
}

var fiCiPrivateCallConfirmed = fieldInfo{
	fType:     FtCiPrivateCallConfirmed,
	typeName:  "Private Call Confimed",
	max:       1,
	bitOffset: 17,
	bitSize:   1,
	valueType: VtOffOn,
	enabler:   FtCiChannelMode,
}

var fiCiQtReverse = fieldInfo{
	fType:     FtCiQtReverse,
	typeName:  "QT Reverse",
	max:       1,
	bitOffset: 36,
	bitSize:   1,
	valueType: VtIStrings,
	strings: &[]string{
		"180",
		"120",
	},
	disabler: FtCiCtcssEncode,
}

var fiCiRepeaterSlot = fieldInfo{
	fType:        FtCiRepeaterSlot,
	typeName:     "Repeater Slot",
	max:          1,
	bitOffset:    12,
	bitSize:      2,
	valueType:    VtIStrings,
	defaultValue: "1",
	strings: &[]string{
		"",
		"1",
		"2",
	},
	enabler: FtCiChannelMode,
}

var fiCiReverseBurst = fieldInfo{
	fType:     FtCiReverseBurst,
	typeName:  "Reverse Burst/Turn Off Code",
	max:       1,
	bitOffset: 37,
	bitSize:   1,
	valueType: VtOffOn,
	disabler:  FtCiCtcssEncode,
}

var fiCiRxFrequency = fieldInfo{
	fType:     FtCiRxFrequency,
	typeName:  "Rx Frequency (MHz)",
	max:       1,
	bitOffset: 128,
	bitSize:   32,
	valueType: VtFrequency,
}

var fiCiRxOnly = fieldInfo{
	fType:     FtCiRxOnly,
	typeName:  "Rx Only",
	max:       1,
	bitOffset: 14,
	bitSize:   1,
	valueType: VtOffOn,
}

var fiCiRxRefFrequency = fieldInfo{
	fType:     FtCiRxRefFrequency,
	typeName:  "Rx Ref Frequency",
	max:       1,
	bitOffset: 30,
	bitSize:   2,
	valueType: VtIStrings,
	strings: &[]string{
		"Low",
		"Medium",
		"High",
	},
}

var fiCiRxSignallingSystem = fieldInfo{
	fType:        FtCiRxSignallingSystem,
	typeName:     "Rx Signaling System",
	max:          1,
	bitOffset:    229,
	bitSize:      3,
	valueType:    VtIStrings,
	defaultValue: "Off",
	strings: &[]string{
		"Off",
		"DTMF-1",
		"DTMF-2",
		"DTMF-3",
		"DTMF-4",
	},
	enablingValue: "Off",
	disabler:      FtCiChannelMode,
}

var fiCiScanList = fieldInfo{
	fType:     FtCiScanList,
	typeName:  "Scan List",
	max:       1,
	bitOffset: 88,
	bitSize:   8,
	valueType: VtListIndex,
	indexedStrings: &[]IndexedString{
		IndexedString{0, "None"},
	},
	listRecordType: RtScanList,
}

var fiCiSquelch = fieldInfo{
	fType:     FtCiSquelch,
	typeName:  "Squelch",
	max:       1,
	bitOffset: 2,
	bitSize:   1,
	valueType: VtIStrings,
	strings: &[]string{
		"Tight",
		"Normal",
	},
}

var fiCiTot = fieldInfo{
	fType:     FtCiTot,
	typeName:  "TOT (S)",
	max:       1,
	bitOffset: 66,
	bitSize:   6,
	valueType: VtSpan,
	span: &Span{
		min:       0,
		max:       63,
		scale:     15,
		interval:  1,
		minString: "Infinite",
	},
}

var fiCiTotRekeyDelay = fieldInfo{
	fType:     FtCiTotRekeyDelay,
	typeName:  "TOT Rekey Delay (S)",
	max:       1,
	bitOffset: 72,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min: 0,
		max: 255,
	},
}

var fiCiTxFrequency = fieldInfo{
	fType:     FtCiTxFrequency,
	typeName:  "Tx Frequency (MHz)",
	max:       1,
	bitOffset: 160,
	bitSize:   32,
	valueType: VtFrequency,
}

var fiCiTxRefFrequency = fieldInfo{
	fType:     FtCiTxRefFrequency,
	typeName:  "Tx Ref Frequency",
	max:       1,
	bitOffset: 38,
	bitSize:   2,
	valueType: VtIStrings,
	strings: &[]string{
		"Low",
		"Medium",
		"High",
	},
}

var fiCiTxSignallingSystem = fieldInfo{
	fType:        FtCiTxSignallingSystem,
	typeName:     "Tx Signaling System",
	max:          1,
	bitOffset:    237,
	bitSize:      3,
	valueType:    VtIStrings,
	defaultValue: "Off",
	strings: &[]string{
		"Off",
		"DTMF-1",
		"DTMF-2",
		"DTMF-3",
		"DTMF-4",
	},
	disabler: FtCiChannelMode,
}

var fiCiVox = fieldInfo{
	fType:     FtCiVox,
	typeName:  "VOX",
	max:       1,
	bitOffset: 35,
	bitSize:   1,
	valueType: VtOffOn,
}

var fiDcCallID = fieldInfo{
	fType:     FtDcCallID,
	typeName:  "Call ID",
	max:       1,
	bitOffset: 0,
	bitSize:   24,
	valueType: VtCallID,
}

var fiDcCallReceiveTone = fieldInfo{
	fType:     FtDcCallReceiveTone,
	typeName:  "Call Receive Tone",
	max:       1,
	bitOffset: 26,
	bitSize:   1,
	valueType: VtIStrings,
	strings: &[]string{
		"No",
		"Yes",
	},
}

var fiDcCallType = fieldInfo{
	fType:     FtDcCallType,
	typeName:  "Call Type",
	max:       1,
	bitOffset: 30,
	bitSize:   2,
	valueType: VtIStrings,
	strings: &[]string{
		"",
		"Group",
		"Private",
		"All",
	},
}

var fiDcContactName = fieldInfo{
	fType:     FtDcContactName,
	typeName:  "Contact Name",
	max:       1,
	bitOffset: 32,
	bitSize:   256,
	valueType: VtName,
}

var fiGlContactMember = fieldInfo{
	fType:          FtGlContactMember,
	typeName:       "Contact Member",
	max:            32,
	bitOffset:      256,
	bitSize:        16,
	valueType:      VtListIndex,
	listRecordType: RtDigitalContacts,
}

var fiGlName = fieldInfo{
	fType:     FtGlName,
	typeName:  "Group List Name",
	max:       1,
	bitOffset: 0,
	bitSize:   256,
	valueType: VtName,
}

var fiGsBacklightColor = fieldInfo{
	fType:     FtGsBacklightColor,
	typeName:  "Backlight Color",
	max:       1,
	bitOffset: 542,
	bitSize:   2,
	valueType: VtIStrings,
	strings: &[]string{
		"Off",
		"Orange",
		"White",
		"Sakura",
	},
}

var fiGsBacklightTime = fieldInfo{
	fType:     FtGsBacklightTime,
	typeName:  "Backlight Time (S)",
	max:       1,
	bitOffset: 686,
	bitSize:   2,
	valueType: VtIStrings,
	strings: &[]string{
		"Always",
		"5",
		"10",
		"15",
	},
}

var fiGsCallAlertToneDuration = fieldInfo{
	fType:     FtGsCallAlertToneDuration,
	typeName:  "Call Alert Tone Duration (S)",
	max:       1,
	bitOffset: 632,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min:       0,
		max:       240,
		scale:     5,
		interval:  1,
		minString: "Continue",
	},
}

var fiGsChFreeIndicationTone = fieldInfo{
	fType:     FtGsChFreeIndicationTone,
	typeName:  "Channel Free Indication Tone",
	max:       1,
	bitOffset: 523,
	bitSize:   1,
	valueType: VtOnOff,
}

var fiGsDisableAllLeds = fieldInfo{
	fType:     FtGsDisableAllLeds,
	typeName:  "Disable All LEDS",
	max:       1,
	bitOffset: 517,
	bitSize:   1,
	valueType: VtOnOff,
}

var fiGsDisableAllTones = fieldInfo{
	fType:     FtGsDisableAllTones,
	typeName:  "Disable All Tones",
	max:       1,
	bitOffset: 525,
	bitSize:   1,
	valueType: VtOnOff,
}

var fiGsFreqChannelMode = fieldInfo{
	fType:     FtGsFreqChannelMode,
	typeName:  "Freq/Channel Mode",
	max:       1,
	bitOffset: 540,
	bitSize:   1,
	valueType: VtIStrings,
	strings: &[]string{
		"Frequency",
		"Channel",
	},
	enablingValue: "Frequency",
}

var fiGsGroupCallHangTime = fieldInfo{
	fType:     FtGsGroupCallHangTime,
	typeName:  "Group Call Hang Time (mS)",
	max:       1,
	bitOffset: 584,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min:      0,
		max:      70,
		scale:    100,
		interval: 5,
	},
}

var fiGsIntroScreen = fieldInfo{
	fType:     FtGsIntroScreen,
	typeName:  "Intro Screen",
	max:       1,
	bitOffset: 531,
	bitSize:   1,
	valueType: VtIStrings,
	strings: &[]string{
		"Character String",
		"Picture",
	},
}

var fiGsIntroScreenLine1 = fieldInfo{
	fType:     FtGsIntroScreenLine1,
	typeName:  "Intro Screen Line 1",
	max:       1,
	bitOffset: 0,
	bitSize:   160,
	valueType: VtIntroLine,
}

var fiGsIntroScreenLine2 = fieldInfo{
	fType:     FtGsIntroScreenLine2,
	typeName:  "Intro Screen Line 2",
	max:       1,
	bitOffset: 160,
	bitSize:   160,
	valueType: VtIntroLine,
}

var fiGsLockUnlock = fieldInfo{
	fType:     FtGsLockUnlock,
	typeName:  "Lock/Unlock",
	max:       1,
	bitOffset: 539,
	bitSize:   1,
	valueType: VtIStrings,
	strings: &[]string{
		"Unlock",
		"Lock",
	},
	disabler: FtGsFreqChannelMode,
}

var fiGsLoneWorkerReminderTime = fieldInfo{
	fType:     FtGsLoneWorkerReminderTime,
	typeName:  "Lone Worker Reminder Time (S)",
	max:       1,
	bitOffset: 648,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min: 1,
		max: 255,
	},
}

var fiGsLoneWorkerResponseTime = fieldInfo{
	fType:     FtGsLoneWorkerResponseTime,
	typeName:  "Lone Worker Response Time (min)",
	max:       1,
	bitOffset: 640,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min: 1,
		max: 255,
	},
}

var fiGsMode = fieldInfo{
	fType:     FtGsMode,
	typeName:  "Mode",
	max:       1,
	bitOffset: 696,
	bitSize:   8,
	valueType: VtIndexedStrings,
	indexedStrings: &[]IndexedString{
		IndexedString{0, "Memory"},
		IndexedString{255, "Channel"},
	},
}

var fiGsModeSelect = fieldInfo{
	fType:     FtGsModeSelect,
	typeName:  "Mode Select",
	max:       1,
	bitOffset: 541,
	bitSize:   1,
	valueType: VtIStrings,
	strings: &[]string{
		"VFO",
		"Memory",
	},
	enabler: FtGsFreqChannelMode,
}

var fiGsMonitorType = fieldInfo{
	fType:     FtGsMonitorType,
	typeName:  "Monitor Type",
	max:       1,
	bitOffset: 515,
	bitSize:   1,
	valueType: VtIStrings,
	strings: &[]string{
		"Silent",
		"Open Squelch",
	},
}

var fiGsPcProgPw = fieldInfo{
	fType:     FtGsPcProgPw,
	typeName:  "PC Programming Password",
	max:       1,
	bitOffset: 768,
	bitSize:   64,
	valueType: VtPcPassword,
}

var fiGsPowerOnPassword = fieldInfo{
	fType:        FtGsPowerOnPassword,
	typeName:     "Power On Password",
	max:          1,
	bitOffset:    704,
	bitSize:      32,
	valueType:    VtRadioPassword,
	defaultValue: "00000000",
	enabler:      FtGsPwAndLockEnable,
}

var fiGsPrivateCallHangTime = fieldInfo{
	fType:     FtGsPrivateCallHangTime,
	typeName:  "Private Call Hang Time (mS)",
	max:       1,
	bitOffset: 592,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min:      0,
		max:      70,
		scale:    100,
		interval: 5,
	},
}

var fiGsPwAndLockEnable = fieldInfo{
	fType:         FtGsPwAndLockEnable,
	typeName:      "Password And Lock Enable",
	max:           1,
	bitOffset:     522,
	bitSize:       1,
	valueType:     VtOnOff,
	enablingValue: "On",
}

var fiGsRadioID = fieldInfo{
	fType:     FtGsRadioID,
	typeName:  "Radio ID",
	max:       1,
	bitOffset: 544,
	bitSize:   24,
	valueType: VtCallID,
}

var fiGsRadioName = fieldInfo{
	fType:     FtGsRadioName,
	typeName:  "Radio Name",
	max:       1,
	bitOffset: 896,
	bitSize:   256,
	valueType: VtRadioName,
}

var fiGsRadioProgPw = fieldInfo{
	fType:     FtGsRadioProgPw,
	typeName:  "Radio Programming Password",
	max:       1,
	bitOffset: 736,
	bitSize:   32,
	valueType: VtRadioPassword,
}

var fiGsRxLowBatteryInterval = fieldInfo{
	fType:     FtGsRxLowBatteryInterval,
	typeName:  "Rx Low Battery Interval (S)",
	max:       1,
	bitOffset: 624,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min:   0,
		max:   127,
		scale: 5,
	},
}

var fiGsSaveModeReceive = fieldInfo{
	fType:     FtGsSaveModeReceive,
	typeName:  "Save Mode Receive",
	max:       1,
	bitOffset: 526,
	bitSize:   1,
	valueType: VtOffOn,
}

var fiGsSavePreamble = fieldInfo{
	fType:     FtGsSavePreamble,
	typeName:  "Save Preamble",
	max:       1,
	bitOffset: 527,
	bitSize:   1,
	valueType: VtOffOn,
}

var fiGsScanAnalogHangTime = fieldInfo{
	fType:     FtGsScanAnalogHangTime,
	typeName:  "Scan Analog Hang Time (mS)",
	max:       1,
	bitOffset: 672,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min:      5,
		max:      100,
		scale:    100,
		interval: 5,
	},
}

var fiGsScanDigitalHangTime = fieldInfo{
	fType:     FtGsScanDigitalHangTime,
	typeName:  "Scan Digital Hang Time (mS)",
	max:       1,
	bitOffset: 664,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min:      5,
		max:      100,
		scale:    100,
		interval: 5,
	},
}

var fiGsSetKeypadLockTime = fieldInfo{
	fType:     FtGsSetKeypadLockTime,
	typeName:  "Set Keypad Lock Time (S)",
	max:       1,
	bitOffset: 688,
	bitSize:   8,
	valueType: VtIndexedStrings,
	indexedStrings: &[]IndexedString{
		IndexedString{255, "Manual"},
		IndexedString{5, "5"},
		IndexedString{10, "10"},
		IndexedString{15, "15"},
	},
}

var fiGsTalkPermitTone = fieldInfo{
	fType:     FtGsTalkPermitTone,
	typeName:  "Talk Permit Tone",
	max:       1,
	bitOffset: 520,
	bitSize:   2,
	valueType: VtIStrings,
	strings: &[]string{
		"None",
		"Digital",
		"Analog",
		"Digital and Analog",
	},
}

var fiGsTxPreambleDuration = fieldInfo{
	fType:     FtGsTxPreambleDuration,
	typeName:  "Tx Preamble Duration (mS)",
	max:       1,
	bitOffset: 576,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min:   0,
		max:   144,
		scale: 60,
	},
}

var fiGsVoxSensitivity = fieldInfo{
	fType:     FtGsVoxSensitivity,
	typeName:  "VOX Sensitivity",
	max:       1,
	bitOffset: 600,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min: 1,
		max: 10,
	},
}

var fiRhHighFrequency = fieldInfo{
	fType:     FtRhHighFrequency,
	typeName:  "High Frequency",
	max:       1,
	bitOffset: 2520,
	bitSize:   16,
	valueType: VtRhFrequency,
}

var fiRhLabel = fieldInfo{
	fType:     FtRhLabel,
	typeName:  "Codeplug Label",
	max:       1,
	bitOffset: 2344,
	bitSize:   128,
	valueType: VtAscii,
}

var fiRhLowFrequency = fieldInfo{
	fType:     FtRhLowFrequency,
	typeName:  "Low Frequency",
	max:       1,
	bitOffset: 2504,
	bitSize:   16,
	valueType: VtRhFrequency,
}

var fiSlChannelMember = fieldInfo{
	fType:          FtSlChannelMember,
	typeName:       "Channel Member",
	max:            31,
	bitOffset:      336,
	bitSize:        16,
	valueType:      VtListIndex,
	listRecordType: RtChannelInformation,
}

var fiSlName = fieldInfo{
	fType:     FtSlName,
	typeName:  "Scan List Name",
	max:       1,
	bitOffset: 0,
	bitSize:   256,
	valueType: VtName,
}

var fiSlPriorityChannel1 = fieldInfo{
	fType:     FtSlPriorityChannel1,
	typeName:  "Priority Channel 1",
	max:       1,
	bitOffset: 256,
	bitSize:   16,
	valueType: VtMemberListIndex,
	indexedStrings: &[]IndexedString{
		IndexedString{0, "Selected"},
		IndexedString{65535, "None"},
	},
	listRecordType: RtChannelInformation,
	enablingValue:  "None",
}

var fiSlPriorityChannel2 = fieldInfo{
	fType:        FtSlPriorityChannel2,
	typeName:     "Priority Channel 2",
	max:          1,
	bitOffset:    272,
	bitSize:      16,
	valueType:    VtMemberListIndex,
	defaultValue: "None",
	indexedStrings: &[]IndexedString{
		IndexedString{0, "Selected"},
		IndexedString{65535, "None"},
	},
	listRecordType: RtChannelInformation,
	disabler:       FtSlPriorityChannel1,
}

var fiSlPrioritySampleTime = fieldInfo{
	fType:     FtSlPrioritySampleTime,
	typeName:  "Priority Sample Time (mS)",
	max:       1,
	bitOffset: 320,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min:   3,
		max:   31,
		scale: 250,
	},
}

var fiSlSignallingHoldTime = fieldInfo{
	fType:     FtSlSignallingHoldTime,
	typeName:  "Signalling Hold Time (mS)",
	max:       1,
	bitOffset: 312,
	bitSize:   8,
	valueType: VtSpan,
	span: &Span{
		min:   2,
		max:   255,
		scale: 25,
	},
}

var fiSlTxDesignatedChannel = fieldInfo{
	fType:     FtSlTxDesignatedChannel,
	typeName:  "Tx Designated Channel",
	max:       1,
	bitOffset: 288,
	bitSize:   16,
	valueType: VtListIndex,
	indexedStrings: &[]IndexedString{
		IndexedString{0, "Selected"},
		IndexedString{65535, "Last Active Channel"},
	},
	listRecordType: RtChannelInformation,
}

var fiTmTextMessage = fieldInfo{
	fType:     FtTmTextMessage,
	typeName:  "Message",
	max:       1,
	bitOffset: 0,
	bitSize:   2304,
	valueType: VtTextMessage,
}

var fiZiChannelMember_md380 = fieldInfo{
	fType:          FtZiChannelMember_md380,
	typeName:       "Channel Member",
	max:            16,
	bitOffset:      256,
	bitSize:        16,
	valueType:      VtListIndex,
	listRecordType: RtChannelInformation,
}

var fiZiChannelMember_md40 = fieldInfo{
	fType:          FtZiChannelMember_md40,
	typeName:       "Channel Member",
	max:            64,
	bitOffset:      256,
	bitSize:        16,
	valueType:      VtListIndex,
	listRecordType: RtChannelInformation,
	extOffset:      201253,
	extSize:        224,
	extIndex:       16,
}

var fiZiName = fieldInfo{
	fType:     FtZiName,
	typeName:  "Zone Name",
	max:       1,
	bitOffset: 0,
	bitSize:   256,
	valueType: VtName,
}

//go:generate genCodeplugInfo
