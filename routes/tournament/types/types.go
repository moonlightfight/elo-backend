package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChallongeBracket struct {
	Tournament struct {
		ID                               int         `json:"id"`
		Name                             string      `json:"name"`
		URL                              string      `json:"url"`
		Description                      string      `json:"description"`
		TournamentType                   string      `json:"tournament_type"`
		StartedAt                        time.Time   `json:"started_at"`
		CompletedAt                      time.Time   `json:"completed_at"`
		RequireScoreAgreement            bool        `json:"require_score_agreement"`
		NotifyUsersWhenMatchesOpen       bool        `json:"notify_users_when_matches_open"`
		CreatedAt                        time.Time   `json:"created_at"`
		UpdatedAt                        time.Time   `json:"updated_at"`
		State                            string      `json:"state"`
		OpenSignup                       bool        `json:"open_signup"`
		NotifyUsersWhenTheTournamentEnds bool        `json:"notify_users_when_the_tournament_ends"`
		ProgressMeter                    int         `json:"progress_meter"`
		QuickAdvance                     bool        `json:"quick_advance"`
		HoldThirdPlaceMatch              bool        `json:"hold_third_place_match"`
		PtsForGameWin                    string      `json:"pts_for_game_win"`
		PtsForGameTie                    string      `json:"pts_for_game_tie"`
		PtsForMatchWin                   string      `json:"pts_for_match_win"`
		PtsForMatchTie                   string      `json:"pts_for_match_tie"`
		PtsForBye                        string      `json:"pts_for_bye"`
		SwissRounds                      int         `json:"swiss_rounds"`
		Private                          bool        `json:"private"`
		RankedBy                         string      `json:"ranked_by"`
		ShowRounds                       bool        `json:"show_rounds"`
		HideForum                        bool        `json:"hide_forum"`
		SequentialPairings               bool        `json:"sequential_pairings"`
		AcceptAttachments                bool        `json:"accept_attachments"`
		RrPtsForGameWin                  string      `json:"rr_pts_for_game_win"`
		RrPtsForGameTie                  string      `json:"rr_pts_for_game_tie"`
		RrPtsForMatchWin                 string      `json:"rr_pts_for_match_win"`
		RrPtsForMatchTie                 string      `json:"rr_pts_for_match_tie"`
		CreatedByAPI                     bool        `json:"created_by_api"`
		CreditCapped                     bool        `json:"credit_capped"`
		Category                         interface{} `json:"category"`
		HideSeeds                        bool        `json:"hide_seeds"`
		PredictionMethod                 int         `json:"prediction_method"`
		PredictionsOpenedAt              interface{} `json:"predictions_opened_at"`
		AnonymousVoting                  bool        `json:"anonymous_voting"`
		MaxPredictionsPerUser            int         `json:"max_predictions_per_user"`
		SignupCap                        interface{} `json:"signup_cap"`
		GameID                           interface{} `json:"game_id"`
		ParticipantsCount                int         `json:"participants_count"`
		GroupStagesEnabled               bool        `json:"group_stages_enabled"`
		AllowParticipantMatchReporting   bool        `json:"allow_participant_match_reporting"`
		Teams                            bool        `json:"teams"`
		CheckInDuration                  interface{} `json:"check_in_duration"`
		StartAt                          interface{} `json:"start_at"`
		StartedCheckingInAt              interface{} `json:"started_checking_in_at"`
		TieBreaks                        []string    `json:"tie_breaks"`
		LockedAt                         interface{} `json:"locked_at"`
		EventID                          interface{} `json:"event_id"`
		PublicPredictionsBeforeStartTime bool        `json:"public_predictions_before_start_time"`
		Ranked                           bool        `json:"ranked"`
		GrandFinalsModifier              interface{} `json:"grand_finals_modifier"`
		PredictTheLosersBracket          bool        `json:"predict_the_losers_bracket"`
		Spam                             interface{} `json:"spam"`
		Ham                              interface{} `json:"ham"`
		RrIterations                     int         `json:"rr_iterations"`
		TournamentRegistrationID         interface{} `json:"tournament_registration_id"`
		DonationContestEnabled           interface{} `json:"donation_contest_enabled"`
		MandatoryDonation                interface{} `json:"mandatory_donation"`
		NonEliminationTournamentData     struct {
		} `json:"non_elimination_tournament_data"`
		AutoAssignStations              interface{} `json:"auto_assign_stations"`
		OnlyStartMatchesWithStations    interface{} `json:"only_start_matches_with_stations"`
		RegistrationFee                 string      `json:"registration_fee"`
		RegistrationType                string      `json:"registration_type"`
		SplitParticipants               bool        `json:"split_participants"`
		AllowedRegions                  interface{} `json:"allowed_regions"`
		ShowParticipantCountry          interface{} `json:"show_participant_country"`
		ProgramID                       interface{} `json:"program_id"`
		ProgramClassificationIdsAllowed interface{} `json:"program_classification_ids_allowed"`
		TeamSizeRange                   interface{} `json:"team_size_range"`
		Toxic                           interface{} `json:"toxic"`
		UseNewStyle                     interface{} `json:"use_new_style"`
		OptionalDisplayData             interface{} `json:"optional_display_data"`
		Participants                    []struct {
			Participant struct {
				ID                                    int           `json:"id"`
				TournamentID                          int           `json:"tournament_id"`
				Name                                  string        `json:"name"`
				Seed                                  int           `json:"seed"`
				Active                                bool          `json:"active"`
				CreatedAt                             time.Time     `json:"created_at"`
				UpdatedAt                             time.Time     `json:"updated_at"`
				InviteEmail                           interface{}   `json:"invite_email"`
				FinalRank                             int           `json:"final_rank"`
				Misc                                  interface{}   `json:"misc"`
				Icon                                  interface{}   `json:"icon"`
				OnWaitingList                         bool          `json:"on_waiting_list"`
				InvitationID                          interface{}   `json:"invitation_id"`
				GroupID                               interface{}   `json:"group_id"`
				CheckedInAt                           interface{}   `json:"checked_in_at"`
				RankedMemberID                        interface{}   `json:"ranked_member_id"`
				CustomFieldResponse                   interface{}   `json:"custom_field_response"`
				Clinch                                interface{}   `json:"clinch"`
				IntegrationUids                       interface{}   `json:"integration_uids"`
				ChallongeUsername                     interface{}   `json:"challonge_username"`
				ChallongeEmailAddressVerified         interface{}   `json:"challonge_email_address_verified"`
				Removable                             bool          `json:"removable"`
				ParticipatableOrInvitationAttached    bool          `json:"participatable_or_invitation_attached"`
				ConfirmRemove                         bool          `json:"confirm_remove"`
				InvitationPending                     bool          `json:"invitation_pending"`
				DisplayNameWithInvitationEmailAddress string        `json:"display_name_with_invitation_email_address"`
				EmailHash                             interface{}   `json:"email_hash"`
				Username                              interface{}   `json:"username"`
				DisplayName                           string        `json:"display_name"`
				AttachedParticipatablePortraitURL     interface{}   `json:"attached_participatable_portrait_url"`
				CanCheckIn                            bool          `json:"can_check_in"`
				CheckedIn                             bool          `json:"checked_in"`
				Reactivatable                         bool          `json:"reactivatable"`
				CheckInOpen                           bool          `json:"check_in_open"`
				GroupPlayerIds                        []interface{} `json:"group_player_ids"`
				HasIrrelevantSeed                     bool          `json:"has_irrelevant_seed"`
			} `json:"participant"`
		} `json:"participants"`
		Matches []struct {
			Match struct {
				ID                        int         `json:"id"`
				TournamentID              int         `json:"tournament_id"`
				State                     string      `json:"state"`
				Player1ID                 int         `json:"player1_id"`
				Player2ID                 int         `json:"player2_id"`
				Player1PrereqMatchID      interface{} `json:"player1_prereq_match_id"`
				Player2PrereqMatchID      interface{} `json:"player2_prereq_match_id"`
				Player1IsPrereqMatchLoser bool        `json:"player1_is_prereq_match_loser"`
				Player2IsPrereqMatchLoser bool        `json:"player2_is_prereq_match_loser"`
				WinnerID                  int         `json:"winner_id"`
				LoserID                   int         `json:"loser_id"`
				StartedAt                 time.Time   `json:"started_at"`
				CreatedAt                 time.Time   `json:"created_at"`
				UpdatedAt                 time.Time   `json:"updated_at"`
				Identifier                string      `json:"identifier"`
				HasAttachment             bool        `json:"has_attachment"`
				Round                     int         `json:"round"`
				Player1Votes              interface{} `json:"player1_votes"`
				Player2Votes              interface{} `json:"player2_votes"`
				GroupID                   interface{} `json:"group_id"`
				AttachmentCount           interface{} `json:"attachment_count"`
				ScheduledTime             interface{} `json:"scheduled_time"`
				Location                  interface{} `json:"location"`
				UnderwayAt                interface{} `json:"underway_at"`
				Optional                  bool        `json:"optional"`
				RushbID                   interface{} `json:"rushb_id"`
				CompletedAt               time.Time   `json:"completed_at"`
				SuggestedPlayOrder        int         `json:"suggested_play_order"`
				Forfeited                 interface{} `json:"forfeited"`
				OpenGraphImageFileName    interface{} `json:"open_graph_image_file_name"`
				OpenGraphImageContentType interface{} `json:"open_graph_image_content_type"`
				OpenGraphImageFileSize    interface{} `json:"open_graph_image_file_size"`
				PrerequisiteMatchIdsCsv   string      `json:"prerequisite_match_ids_csv"`
				ScoresCsv                 string      `json:"scores_csv"`
			} `json:"match"`
		} `json:"matches"`
		DescriptionSource      string      `json:"description_source"`
		Subdomain              interface{} `json:"subdomain"`
		FullChallongeURL       string      `json:"full_challonge_url"`
		LiveImageURL           string      `json:"live_image_url"`
		SignUpURL              interface{} `json:"sign_up_url"`
		ReviewBeforeFinalizing bool        `json:"review_before_finalizing"`
		AcceptingPredictions   bool        `json:"accepting_predictions"`
		ParticipantsLocked     bool        `json:"participants_locked"`
		GameName               interface{} `json:"game_name"`
		ParticipantsSwappable  bool        `json:"participants_swappable"`
		TeamConvertable        bool        `json:"team_convertable"`
		GroupStagesWereStarted bool        `json:"group_stages_were_started"`
	} `json:"tournament"`
}

type BracketInfo struct {
	Title          string    `json:"title"`
	NumPlayers     int       `json:"numPlayers"`
	TournamentDate time.Time `json:"tournamentDate"`
	Players        []Player  `json:"players"`
	Matches        []Match   `json:"matches"`
}

type Match struct {
	WinnerID    interface{} `json:"winnerId"`
	LoserID     interface{} `json:"loserId"`
	WinnerName  string      `json:"winnerName"`
	LoserName   string      `json:"loserName"`
	WinnerScore int         `json:"winnerScore"`
	LoserScore  int         `json:"loserScore"`
	MatchDate   time.Time   `json:"matchTime"`
}

type ReturnedData struct {
	Token      string `json:"token"`
	Tournament struct {
		*BracketInfo
		Replay   string `json:"replay"`
		Location string `json:"location"`
	} `json:"tournament"`
}

type Player struct {
	ID    primitive.ObjectID `json:"id"`
	Name  string             `json:"name"`
	Place int                `json:"place"`
}

type SmashQuery struct {
	Query     string         `json:"query"`
	Variables SmashVariables `json:"variables"`
}

type SmashVariables struct {
	Slug string `json:"slug"`
}

type SmashBracket struct {
	Data struct {
		Event struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			StartAt   int64  `json:"startAt"`
			Standings struct {
				Nodes []struct {
					ID        int `json:"id"`
					Placement int `json:"placement"`
					Entrant   struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"entrant"`
				} `json:"nodes"`
			} `json:"standings"`
			Sets struct {
				Nodes []struct {
					ID    int `json:"id"`
					Slots []struct {
						Entrant struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
						} `json:"entrant"`
					} `json:"slots"`
					Winnerid     int    `json:"winnerId"`
					Displayscore string `json:"displayScore"`
					CompletedAt  int64  `json:"completedAt"`
				} `json:"nodes"`
			} `json:"sets"`
			Videogame struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"videogame"`
			Tournament struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"tournament"`
		} `json:"event"`
	} `json:"data"`
	Extensions struct {
		Cachecontrol struct {
			Version int `json:"version"`
			Hints   []struct {
				Path   []string `json:"path"`
				Maxage int      `json:"maxAge"`
				Scope  string   `json:"scope"`
			} `json:"hints"`
		} `json:"cacheControl"`
		Querycomplexity int `json:"queryComplexity"`
	} `json:"extensions"`
	Actionrecords []interface{} `json:"actionRecords"`
}
