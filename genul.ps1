go build

#golf data feeds
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_currenttournaments.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Feed_CurrentTournament -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_tournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Feed_Tournament -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_teamsummary.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Feed_TeamSummary -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_teammatchsummary.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Feed_TeamMatchSummary -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_teammatchscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Feed_TeamMatchScorecard -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_courses.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Feed_Courses -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_golfdatafeed_playoffs.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Feed_Playoffs -ns LPGALiveScoringServices.ULCrown


#tournament
.\lpgagen -file datafiles\ulcrown\ul_tournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Tournament -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_tournamentround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TournamentRound -ns LPGALiveScoringServices.ULCrown

#team summary

.\lpgagen -file datafiles\ulcrown\ul_teamsummarytournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamSummaryTournament -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teamsummarytournamentround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamSummaryTournamentRound -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teamsummaryroundpoints.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamSummaryRoundPoints -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teamsummary.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamSummary -ns LPGALiveScoringServices.ULCrown

#team match summary
.\lpgagen -file datafiles\ulcrown\ul_teammatchsummarytournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamMatchSummaryTournament -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchsummarytournamentround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamMatchSummaryTournamentRound -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchsummarymatch.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamMatchSummaryMatch -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchsummaryteam.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamMatchSummaryTeam -ns LPGALiveScoringServices.ULCrown

#team match scorecard
.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardtournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamMatchScorecardTournament -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamMatchScorecardRound -ns LPGALiveScoringServices.ULCrown

.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardmatch.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamMatchScorecardMatch -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardteam.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamMatchScorecardTeam -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardteamscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamMatchScorecardTeamScorecard -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_teammatchscorecardteamscorecardscore.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TeamMatchScorecardTeamScorecardScore -ns LPGALiveScoringServices.ULCrown

#course
.\lpgagen -file datafiles\ulcrown\ul_coursetournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_CourseTournament -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_course.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Course -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_coursehole.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_CourseHole -ns LPGALiveScoringServices.ULCrown


#playoffs
.\lpgagen -file datafiles\ulcrown\ul_playoff.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Playoff -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffcourse.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayoffCourse -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffcoursehole.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayoffCourseHole -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffplayerscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayoffPlayerScorecard -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffplayerscorecardscore.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayoffPlayerScorecardScore -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffteam.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayoffTeam -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffplayer.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayoffPlayer -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffteamscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayoffTeamScorecard -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playoffteamscorecardscore.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayoffTeamScorecardScore -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playofftournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayoffTournament -ns LPGALiveScoringServices.ULCrown


#player

.\lpgagen -file datafiles\ulcrown\ul_player.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Player -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playerscorecard.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayerScorecard -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ulcrown\ul_playerscorecardscore.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_PlayerScorecardScore -ns LPGALiveScoringServices.ULCrown
