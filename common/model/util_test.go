package model

import (
	"testing"
	"log"
)

func TestTrimAll(t *testing.T) {
	log.Println(string(text))
	log.Println(TrimAll(string(text)))
}

func BenchmarkTrimAll(b *testing.B) {


	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TrimAll(string(text))
	}
}

var text []byte = []byte(`[{"  \x00   \x01     subjectNotes":null,"       courseNumber":"101","subject":"812","campusCode":"NK","openSections":1,"synopsisUrl":null,"subjectGroupNotes":null,"offeringUnitCode":"21","offeringUnitTitle":null,"title":"ELEM PORT","courseDescription":null,"preReqNotes":null,"sections":[{"sectionEligibility":null,"sessionDatePrintIndicator":"Y","examCode":"C","specialPermissionAddCode":null,"crossListedSections":[],"sectionNotes":"BEGINNERS ONLY. PLACEMENT EXAM/CHAIR'S PERMISSION REQUIRED FORSTUDENTS WITH PRIOR KNOWLEDGEOF PORTUGUESE.","specialPermissionDropCode":null,"instructors":[{"name":"CASTILHO"}],"number":"01","majors":[],"sessionDates":null,"specialPermissionDropCodeDescription":null,"subtopic":null,"openStatus":true,"comments":[],"minors":[],"campusCode":"NK","index":"09229","unitMajors":[],"printed":"Y","specialPermissionAddCodeDescription":null,"subtitle":null,"meetingTimes":[{"campusLocation":"7","baClassHours":null,"roomNumber":"455","pmCode":"A","campusAbbrev":"NWK","campusName":"NEWARK","meetingDay":"T","buildingCode":"CON","startTime":"1000","endTime":"1120","meetingModeDesc":"LEC","meetingModeCode":"02"},{"campusLocation":"7","baClassHours":null,"roomNumber":"455","pmCode":"A","campusAbbrev":"NWK","campusName":"NEWARK","meetingDay":"TH","buildingCode":"CON","startTime":"1000","endTime":"1120","meetingModeDesc":"LEC","meetingModeCode":"02"}],"legendKey":null,"honorPrograms":[]}],"supplementCode":" ","credits":3,"unitNotes":null,"coreCodes":[],"courseNotes":null,"expandedTitle":"ELEMENTARY PORTUGUESE "},{"subjectNotes":null,"courseNumber":"131","subject":"812","campusCode":"NK","openSections":1,"synopsisUrl":null,"subjectGroupNotes":null,"offeringUnitCode":"21","offeringUnitTitle":null,"title":"INTER PORTUGUESE","courseDescription":null,"preReqNotes":"(21:812:102 ) OR (21:810:102 ELEM PORT)","sections":[{"sectionEligibility":null,"sessionDatePrintIndicator":"Y","examCode":"C","specialPermissionAddCode":null,"crossListedSections":[],"sectionNotes":"PREREQ: 812:102, 810:102,PLACEMENT EXAM, OR PERMISSION OFINSTRUCTOR.","specialPermissionDropCode":null,"instructors":[{"name":"CASTILHO"}],"number":"01","majors":[],"sessionDates":null,"specialPermissionDropCodeDescription":null,"subtopic":null,"openStatus":true,"comments":[],"minors":[],"campusCode":"NK","index":"09228","unitMajors":[],"printed":"Y","specialPermissionAddCodeDescription":null,"subtitle":null,"meetingTimes":[{"campusLocation":"7","baClassHours":null,"roomNumber":"455","pmCode":"A","campusAbbrev":"NWK","campusName":"NEWARK","meetingDay":"TH","buildingCode":"CON","startTime":"1130","endTime":"1250","meetingModeDesc":"LEC","meetingModeCode":"02"},{"campusLocation":"7","baClassHours":null,"roomNumber":"455","pmCode":"A","campusAbbrev":"NWK","campusName":"NEWARK","meetingDay":"T","buildingCode":"CON","startTime":"1130","endTime":"1250","meetingModeDesc":"LEC","meetingModeCode":"02"}],"legendKey":null,"honorPrograms":[]}],"supplementCode":" ","credits":3,"unitNotes":null,"coreCodes":[],"courseNotes":null,"expandedTitle":"INTERMEDIATE PORTUGUESE "},{"subjectNotes":null,"courseNumber":"203","subject":"812","campusCode":"NK","openSections":1,"synopsisUrl":null,"subjectGroupNotes":null,"offeringUnitCode":"21","offeringUnitTitle":null,"title":"ADV GRAMMAR & COMP","courseDescription":null,"preReqNotes":"((21:355:102 or 62:355:102 ENGLISH COMPOSITION) and (21:812:132 )) OR ((21:355:104 or 62:355:104 HONORS ENGLISH COMP) and (21:812:132 )) OR ((21:355:102 or 62:355:102 ENGLISH COMPOSITION) and (21:810:132 INTER PORTUGUESE)) OR ((21:355:104 or 62:355:104 HONORS ENGLISH COMP) and (21:810:132 INTER PORTUGUESE))","sections":[{"sectionEligibility":null,"sessionDatePrintIndicator":"Y","examCode":"C","specialPermissionAddCode":null,"crossListedSections":[],"sectionNotes":null,"specialPermissionDropCode":null,"instructors":[{"name":"CASTILHO"}],"number":"01","majors":[],"sessionDates":null,"specialPermissionDropCodeDescription":null,"subtopic":null,"openStatus":true,"comments":[],"minors":[],"campusCode":"NK","index":"16569","unitMajors":[],"printed":"Y","specialPermissionAddCodeDescription":null,"subtitle":null,"meetingTimes":[{"campusLocation":"7","baClassHours":null,"roomNumber":"348","pmCode":"P","campusAbbrev":"NWK","campusName":"NEWARK","meetingDay":"T","buildingCode":"CON","startTime":"0230","endTime":"0350","meetingModeDesc":"LEC","meetingModeCode":"02"},{"campusLocation":"7","baClassHours":null,"roomNumber":"348","pmCode":"P","campusAbbrev":"NWK","campusName":"NEWARK","meetingDay":"TH","buildingCode":"CON","startTime":"0230","endTime":"0350","meetingModeDesc":"LEC","meetingModeCode":"02"}],"legendKey":null,"honorPrograms":[]}],"supplementCode":" ","credits":3,"unitNotes":null,"coreCodes":[],"courseNotes":null,"expandedTitle":"ADVANCED GRAMMAR & COMPOSITION "},{"subjectNotes":null,"courseNumber":"250","subject":"812","campusCode":"NK","openSections":0,"synopsisUrl":null,"subjectGroupNotes":null,"offeringUnitCode":"21","offeringUnitTitle":null,"title":"PORT LIT ENG TRAN","courseDescription":null,"preReqNotes":null,"sections":[{"sectionEligibility":null,"sessionDatePrintIndicator":"Y","examCode":"C","specialPermissionAddCode":null,"crossListedSections":[],"sectionNotes":"TAUGHT IN ENGLISH","specialPermissionDropCode":null,"instructors":[{"name":"HOLTON"}],"number":"01","majors":[],"sessionDates":null,"specialPermissionDropCodeDescription":null,"subtopic":null,"openStatus":false,"comments":[{"code":"20","description":"Taught in English"}],"minors":[],"campusCode":"NK","index":"16572","unitMajors":[],"printed":"Y","specialPermissionAddCodeDescription":null,"subtitle":"TALES OF TRAVEL ","meetingTimes":[{"campusLocation":"7","baClassHours":null,"roomNumber":"211","pmCode":"P","campusAbbrev":"NWK","campusName":"NEWARK","meetingDay":"T","buildingCode":"ENG","startTime":"0230","endTime":"0520","meetingModeDesc":"LEC","meetingModeCode":"02"}],"legendKey":null,"honorPrograms":[]}],"supplementCode":" ","credits":3,"unitNotes":null,"coreCodes":[{"id":"2016921812250 31","year":"2016","description":"History & Literature, Literature","effective":"20169","coreCodeReferenceId":"31","offeringUnitCode":"21","coreCode":"HL:Lt","coreCodeDescription":"History & Literature, Literature","lastUpdated":1468423866000,"subject":"812","code":"HL:Lt","term":"9","unit":"21","course":"250","offeringUnitCampus":"NK","supplement":" "}],"courseNotes":null,"expandedTitle":"PORTUGUESE LITERATURE IN ENGLISH TRANSLATION "},{"subjectNotes":null,"courseNumber":"319","subject":"812","campusCode":"NK","openSections":1,"synopsisUrl":null,"subjectGroupNotes":null,"offeringUnitCode":"21","offeringUnitTitle":null,"title":"BRAZILIAN LIT TRANSL","courseDescription":null,"preReqNotes":null,"sections":[{"sectionEligibility":null,"sessionDatePrintIndicator":"Y","examCode":"S","specialPermissionAddCode":null,"crossListedSections":[],"sectionNotes":null,"specialPermissionDropCode":null,"instructors":[{"name":"NASCIMENTO"}],"number":"61","majors":[],"sessionDates":null,"specialPermissionDropCodeDescription":null,"subtopic":null,"openStatus":true,"comments":[],"minors":[],"campusCode":"NK","index":"19875","unitMajors":[],"printed":"Y","specialPermissionAddCodeDescription":null,"subtitle":null,"meetingTimes":[{"campusLocation":"7","baClassHours":null,"roomNumber":"342","pmCode":"P","campusAbbrev":"NWK","campusName":"NEWARK","meetingDay":"T","buildingCode":"CON","startTime":"0600","endTime":"0900","meetingModeDesc":"LEC","meetingModeCode":"02"}],"legendKey":null,"honorPrograms":[]}],"supplementCode":" ","credits":3,"unitNotes":null,"coreCodes":[{"id":"2016921812319 31","year":"2016","description":"History & Literature, Literature","effective":"20169","coreCodeReferenceId":"31","offeringUnitCode":"21","coreCode":"HL:Lt","coreCodeDescription":"History & Literature, Literature","lastUpdated":1468423866000,"subject":"812","code":"HL:Lt","term":"9","unit":"21","course":"319","offeringUnitCampus":"NK","supplement":" "}],"courseNotes":null,"expandedTitle":"BRAZILIAN LITERATURE IN ENGLISH TRANSLATION II "},{"subjectNotes":null,"courseNumber":"423","subject":"812","campusCode":"NK","openSections":1,"synopsisUrl":null,"subjectGroupNotes":null,"offeringUnitCode":"21","offeringUnitTitle":null,"title":"POSTREVPORTLIT&CULT","courseDescription":null,"preReqNotes":null,"sections":[{"sectionEligibility":null,"sessionDatePrintIndicator":"Y","examCode":"C","specialPermissionAddCode":null,"crossListedSections":[],"sectionNotes":"TAUGHT IN PORTUGUESE","specialPermissionDropCode":null,"instructors":[{"name":"HOLTON"}],"number":"01","majors":[],"sessionDates":null,"specialPermissionDropCodeDescription":null,"subtopic":null,"openStatus":true,"comments":[],"minors":[],"campusCode":"NK","index":"18906","unitMajors":[],"printed":"Y","specialPermissionAddCodeDescription":null,"subtitle":null,"meetingTimes":[{"campusLocation":"7","baClassHours":null,"roomNumber":"211","pmCode":"P","campusAbbrev":"NWK","campusName":"NEWARK","meetingDay":"TH","buildingCode":"ENG","startTime":"0230","endTime":"0520","meetingModeDesc":"LEC","meetingModeCode":"02"}],"legendKey":null,"honorPrograms":[]}],"supplementCode":" ","credits":3,"unitNotes":null,"coreCodes":[],"courseNotes":null,"expandedTitle":"POST-REVOLUTIONARY PORTUGUESE LITERATURE& CULTURE "},{"subjectNotes":null,"courseNumber":"458","subject":"812","campusCode":"NK","openSections":1,"synopsisUrl":null,"subjectGroupNotes":null,"offeringUnitCode":"21","offeringUnitTitle":null,"title":"INTERN PORT STUDIES","courseDescription":null,"preReqNotes":null,"sections":[{"sectionEligibility":null,"sessionDatePrintIndicator":"Y","examCode":"A","specialPermissionAddCode":"16","crossListedSections":[],"sectionNotes":"BY PERMISSION OF DEPT CHAIR;FOR JUNIORS & SENIORS MAJORINGOR MINORING IN PORTUGUESE.","specialPermissionDropCode":null,"instructors":[{"name":"STAFF"}],"number":"01","majors":[],"sessionDates":null,"specialPermissionDropCodeDescription":null,"subtopic":null,"openStatus":true,"comments":[],"minors":[],"campusCode":"NK","index":"20109","unitMajors":[],"printed":"Y","specialPermissionAddCodeDescription":"Department","subtitle":null,"meetingTimes":[{"campusLocation":"7","baClassHours":"B","roomNumber":null,"pmCode":null,"campusAbbrev":"NWK","campusName":"NEWARK","meetingDay":null,"buildingCode":null,"startTime":null,"endTime":null,"meetingModeDesc":"INTERNSP","meetingModeCode":"15"}],"legendKey":null,"honorPrograms":[]}],"supplementCode":" ","credits":3,"unitNotes":null,"coreCodes":[],"courseNotes":null,"expandedTitle":"INTERNSHIP IN PORTUGUESE & LUSOPHONE WORLD STUDIES "}]`)