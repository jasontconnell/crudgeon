go build

.\lpgagen -file datafiles\stats\z_careerMoney.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_CareerMoney -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_careerstats.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_CareerStats -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_dataheartbeat.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_DataHeartbeat -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_playertournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_PlayerTournament -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_playeryear.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_PlayerYear -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_ptssched.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_RacePointsSchedule -ns LPGAStatsService.Data -usefield

.\lpgagen -file datafiles\stats\z_cmeplayerseason.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_CMEPlayerSeason -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_cmeplayerfinal.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_CMEPlayerFinal -ns LPGAStatsService.Data -usefield


.\lpgagen -file datafiles\stats\z_rolexmajoraward.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_RolexMajorAward -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_tournaments.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_Tournament -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_tourncourse.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_TournamentCourse -ns LPGAStatsService.Data -usefield

.\lpgagen -file datafiles\stats\z_wwtourndonamt.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_WWTournDonationAmount -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_wwtourndonamtround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_WWTournDonationAmountRound -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_wwytddonamt.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_WWYTDDonationAmount -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_wwytdplayersummary.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_WWYTDPlayerSummary -ns LPGAStatsService.Data -usefield
.\lpgagen -file datafiles\stats\z_wwytdsummary.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_WWYTDSummary -ns LPGAStatsService.Data -usefield


.\lpgagen -file datafiles\stats\api_tournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_Tournament -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_tournamentfield.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentField -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_tournamentfieldplayer.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentFieldPlayer -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_tournamentcourse.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentCourse -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_tournamentcoursehole.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentCourseHole -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_player.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_Player -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_playerscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerScorecard -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_playerscorecardround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerScorecardRound -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_playerscorecardroundhole.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerScorecardRoundHole -ns LPGAStatsService.Data

.\lpgagen -file datafiles\stats\api_playerhistory.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerHistory -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_playerhistoryevent.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerHistoryEvent -ns LPGAStatsService.Data

.\lpgagen -file datafiles\stats\api_playerquickviewofficialmoneycomponents.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewOfficialMoneyComponentList -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_playerquickviewofficialmoney.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewOfficialMoneyComponent -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_playerquickviewofficialmoneystatistic.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewOfficialMoneyComponentStatistic -ns LPGAStatsService.Data

.\lpgagen -file datafiles\stats\api_cmegloberace.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_CMEGlobeRace -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_cmegloberacemember.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_CMEGlobeRaceMember -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_cmegloberacestat.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_CMEGlobeRaceStat -ns LPGAStatsService.Data

.\lpgagen -file datafiles\stats\api_cmepointsschedule.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_CMEPointsSchedule -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_cmepointsscheduleentry.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_CMEPointsScheduleEntry -ns LPGAStatsService.Data

.\lpgagen -file datafiles\stats\api_playerrolexpoints.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerRolexPoints -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_playerrolexpointsevent.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerRolexPointsEvent -ns LPGAStatsService.Data

.\lpgagen -file datafiles\stats\api_tournamentparticipant.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentParticipantList -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_tournamentparticipantentry.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentParticipantListEntry -ns LPGAStatsService.Data

.\lpgagen -file datafiles\stats\api_playerperformance.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerPerformance -ns LPGAStatsService.Data


.\lpgagen -file datafiles\stats\api_tournamenthistory.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentHistory -ns LPGAStatsService.Data

.\lpgagen -file datafiles\stats\api_playerquickviewboxscorelist.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewBoxScoreList -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_playerquickviewboxscorelistentry.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewBoxScoreListEntry -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_playerquickviewboxscorelistentryround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewBoxScoreListEntryRound -ns LPGAStatsService.Data

.\lpgagen -file datafiles\stats\api_tournamentyearlistentry.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentYearListEntry -ns LPGAStatsService.Data

.\lpgagen -file datafiles\stats\api_solheim.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_Solheim -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_solheimmember.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_SolheimMember -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\api_solheimmemberstat.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_SolheimMemberStatistic -ns LPGAStatsService.Data


.\lpgagen -file datafiles\stats\r_rankwrapper.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj RankingAPI_Wrapper -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\r_rankingweek.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj RankingAPI_Week -ns LPGAStatsService.Data
.\lpgagen -file datafiles\stats\r_rankingplayer.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj RankingAPI_Player -ns LPGAStatsService.Data
