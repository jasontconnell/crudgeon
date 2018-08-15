go build

.\lpgagen -file datafiles\z_careerMoney.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_CareerMoney -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_careerstats.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_CareerStats -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_dataheartbeat.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_DataHeartbeat -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_playertournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_PlayerTournament -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_playeryear.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_PlayerYear -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_ptssched.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_RacePointsSchedule -ns LPGAStatsService.Data

.\lpgagen -file datafiles\z_cmeplayerseason.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_CMEPlayerSeason -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_cmeplayerfinal.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_CMEPlayerFinal -ns LPGAStatsService.Data


.\lpgagen -file datafiles\z_rolexmajoraward.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_RolexMajorAward -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_tournaments.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_Tournament -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_tourncourse.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_TournamentCourse -ns LPGAStatsService.Data

.\lpgagen -file datafiles\z_wwtourndonamt.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_WWTournDonationAmount -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_wwtourndonamtround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_WWTournDonationAmountRound -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_wwytddonamt.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_WWYTDDonationAmount -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_wwytdplayersummary.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_WWYTDPlayerSummary -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_wwytdsummary.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_WWYTDSummary -ns LPGAStatsService.Data




.\lpgagen -file datafiles\z_sym_careerstats.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_SymetraCareerStats -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_sym_playertournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_SymetraPlayerTournament -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_sym_playeryear.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_SymetraPlayerYear -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_sym_playertournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_SymetraPlayerTournament -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_sym_tournaments.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_SymetraTournament -ns LPGAStatsService.Data
.\lpgagen -file datafiles\z_sym_tourncourse.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj Stats_SymetraTournamentCourse -ns LPGAStatsService.Data



.\lpgagen -file datafiles\api_tournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_Tournament -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_tournamentfield.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentField -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_tournamentfieldplayer.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentFieldPlayer -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_tournamentcourse.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentCourse -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_tournamentcoursehole.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentCourseHole -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_player.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_Player -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_playerscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerScorecard -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_playerscorecardround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerScorecard -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_playerscorecardroundhole.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerScorecardRoundHole -ns LPGAStatsService.Data

.\lpgagen -file datafiles\api_playerhistory.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerHistory -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_playerhistoryevent.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerHistoryEvent -ns LPGAStatsService.Data

.\lpgagen -file datafiles\api_playerquickviewofficialmoney.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewOfficialMoney -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_playerquickviewofficialmoneystatistic.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewOfficialMoneyStatistic -ns LPGAStatsService.Data

.\lpgagen -file datafiles\api_cmegloberace.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_CMEGlobeRace -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_cmegloberacemember.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_CMEGlobeRaceMember -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_cmegloberacestat.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_CMEGlobeRaceStat -ns LPGAStatsService.Data

.\lpgagen -file datafiles\api_cmepointsschedule.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_CMEPointsSchedule -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_cmepointsscheduleentry.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_CMEPointsScheduleEntry -ns LPGAStatsService.Data

.\lpgagen -file datafiles\api_playerrolexpoints.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerRolexPoints -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_playerrolexpointsevent.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerRolexPointsEvent -ns LPGAStatsService.Data

.\lpgagen -file datafiles\api_tournamentparticipant.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentParticipantList -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_tournamentparticipantentry.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentParticipantListEntry -ns LPGAStatsService.Data

.\lpgagen -file datafiles\api_playerperformance.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerPerformance -ns LPGAStatsService.Data


.\lpgagen -file datafiles\api_tournamenthistory.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentHistory -ns LPGAStatsService.Data

.\lpgagen -file datafiles\api_playerquickviewboxscorelist.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewBoxScoreList -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_playerquickviewboxscorelistentry.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewBoxScoreListEntry -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_playerquickviewboxscorelistentryround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_PlayerQuickViewBoxScoreListEntryRound -ns LPGAStatsService.Data

.\lpgagen -file datafiles\api_tournamentyearlist.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentYearList -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_tournamentyearlistentry.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_TournamentYearListEntry -ns LPGAStatsService.Data

.\lpgagen -file datafiles\api_solheim.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_Solheim -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_solheimmember.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_SolheimMember -ns LPGAStatsService.Data
.\lpgagen -file datafiles\api_solheimmemberstat.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGAStatsService.Data" -obj StatsAPI_SolheimMemberStatistic -ns LPGAStatsService.Data
