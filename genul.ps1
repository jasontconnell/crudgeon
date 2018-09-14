go build

#golf data feeds
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_currenttournaments.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Feed_CurrentTournament -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_tournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Feed_Tournament -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_teamsummary.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Feed_TeamSummary -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_teammatchsummary.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Feed_TeamMatchSummary -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_teammatchscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Feed_TeamMatchScorecard -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_courses.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Feed_Courses -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_playoffs.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Feed_Playoffs -ns LPGALiveScoring.Data.ULCrown


#tournament
.\lpgagen -file datafiles\ulcrown\ul_tournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Tournament -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_tournamentround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TournamentRound -ns LPGALiveScoring.Data.ULCrown

#team summary

.\lpgagen -file datafiles\ulcrown\ul_teamsummarytournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamSummaryTournament -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teamsummarytournamentround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamSummaryTournamentRound -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teamsummaryroundpoints.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamSummaryRoundPoints -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teamsummary.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamSummary -ns LPGALiveScoring.Data.ULCrown

#team match summary
.\lpgagen -file datafiles\ulcrown\ul_teammatchsummarytournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamMatchSummaryTournament -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchsummarytournamentround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamMatchSummaryTournamentRound -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchsummarymatch.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamMatchSummaryMatch -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchsummaryteam.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamMatchSummaryTeam -ns LPGALiveScoring.Data.ULCrown

#team match scorecard
.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardtournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamMatchScorecardTournament -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamMatchScorecardRound -ns LPGALiveScoring.Data.ULCrown

.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardmatch.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamMatchScorecardMatch -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardteam.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamMatchScorecardTeam -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardteamscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamMatchScorecardTeamScorecard -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardteamscorecardscore.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_TeamMatchScorecardTeamScorecardScore -ns LPGALiveScoring.Data.ULCrown

#course
.\lpgagen -file datafiles\ulcrown\ul_coursetournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_CourseTournament -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_course.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Course -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_coursehole.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_CourseHole -ns LPGALiveScoring.Data.ULCrown


#playoffs
.\lpgagen -file datafiles\ulcrown\ul_playoff.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Playoff -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffcourse.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayoffCourse -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffcoursehole.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayoffCourseHole -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffplayerscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayoffPlayerScorecard -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffplayerscorecardscore.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayoffPlayerScorecardScore -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffteam.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayoffTeam -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffplayer.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayoffPlayer -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffteamscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayoffTeamScorecard -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffteamscorecardscore.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayoffTeamScorecardScore -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playofftournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayoffTournament -ns LPGALiveScoring.Data.ULCrown


#player

.\lpgagen -file datafiles\ulcrown\ul_player.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_Player -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playerscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayerScorecard -ns LPGALiveScoring.Data.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playerscorecardscore.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoring.Data\ULCrown" -obj ULCrown_PlayerScorecardScore -ns LPGALiveScoring.Data.ULCrown
