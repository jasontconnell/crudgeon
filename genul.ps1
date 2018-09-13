go build

.\lpgagen -file datafiles\ul_golfdatafeed_currenttournaments.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Feed_CurrentTournament -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ul_golfdatafeed_tournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Feed_Tournament -ns LPGALiveScoringServices.ULCrown

.\lpgagen -file datafiles\ul_tournament.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_Tournament -ns LPGALiveScoringServices.ULCrown
.\lpgagen -file datafiles\ul_tournamentround.txt -path "C:\Users\jconnell\source\repos\LPGALiveScoring\LPGALiveScoringServices\ULCrown" -obj ULCrown_TournamentRound -ns LPGALiveScoringServices.ULCrown
