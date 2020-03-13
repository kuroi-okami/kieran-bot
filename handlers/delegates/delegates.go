package delegates

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const rob string = "293878271844679680"
const will string = "291248097898528778"
const steve string = "284441212297936907"
const clark string = "284441148653830144"

const danish string = "531505474642771986"
const cowboy string = "383338439539687427"
const die string = "471402580497530890"
const join string = "655514010833387577"
const not string = "687809491898859640"
const weague string = "687809343797723211"
const corona string = "687809745213718548"
const pog string = "687809799387349066"
const pepe string = "687809850801258665"
const big string = "687809911719329838"

var channelArr = []string{
	danish,
	cowboy,
	die,
	join,
	not,
	weague,
	corona,
	pog,
	pepe,
	big,
}

var clarkovEnabled bool = false

func KieranSpeak(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!story" {
		rand.Seed(time.Now().UnixNano())
		storyId := rand.Intn(10)
		switch storyId {
		case 1:
			{
				s.ChannelMessageSend(m.ChannelID, "I once fed a homeless man crumbs")
			}
		case 2:
			{
				s.ChannelMessageSend(m.ChannelID, "Once James and Will froze my shoes")
			}
		case 3:
			{
				s.ChannelMessageSend(m.ChannelID, "I love jesus")
			}
		case 4:
			{
				s.ChannelMessageSend(m.ChannelID, "Have you meet jesus?")
			}
		case 5:
			{
				s.ChannelMessageSend(m.ChannelID, "I have cute feet")
			}
		case 6:
			{
				s.ChannelMessageSend(m.ChannelID, "What is love?")
			}
		case 7:
			{
				s.ChannelMessageSend(m.ChannelID, "You guys are fudged in the head")
			}
		case 8:
			{
				s.ChannelMessageSend(m.ChannelID, "I ate a cake that was too terribly good to eat badly")
			}
		case 9:
			{
				s.ChannelMessageSend(m.ChannelID, "I miss limelight girl")
			}
		case 10:
			{
				s.ChannelMessageSend(m.ChannelID, "You guys are lame")
			}
		default:
			{
				s.ChannelMessageSend(m.ChannelID, "I am all out of stories")
			}
		}
	}
}

func RemindClark(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!remindclark") {
		commandStr := strings.TrimPrefix(m.Content, "!remindclark ")

		const delimiter string = " "
		const numberOfArgs int = 2
		arguments := strings.SplitN(commandStr, delimiter, numberOfArgs)
		if len(arguments) != numberOfArgs {
			s.ChannelMessageSend(m.ChannelID, "CLARK. WRONG FUCKING ARGUMENTS")
			return
		}

		duration, err := strconv.Atoi(arguments[0])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "CLARK. THAT'S NOT A FUCKING INT")
			return
		}

		clark, err := s.User("284441148653830144")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error 404: Clark not found")
		}

		time.AfterFunc(
			time.Second * time.Duration(duration),
			func(){
				s.ChannelMessageSend(m.ChannelID, clark.Mention() + " " + arguments[1] )
			})

		s.ChannelMessageSend(m.ChannelID, "Your pingas has been installed...")
	}
}


func LoopSteve(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!loopsteve") {
		s.ChannelMessageSend(m.ChannelID, "Steve no likey so feature has been deprecatedy")
		rob, err := s.User("293878271844679680")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error 404: Rob not found")
		}
		s.ChannelMessageSend(m.ChannelID, "You can still ddos " + rob.Mention() + " with !ddosrob though")
		return

		commandStr := strings.TrimPrefix(m.Content, "!loopsteve ")

		const delimiter string = " "
		const numberOfArgs int = 2
		arguments := strings.SplitN(commandStr, delimiter, numberOfArgs)
		if len(arguments) != numberOfArgs {
			s.ChannelMessageSend(m.ChannelID, "STEVE. WRONG FUCKING ARGUMENTS")
			return
		}

		numLoops, err := strconv.Atoi(arguments[0])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "STEVE. THAT'S NOT A FUCKING INT")
			return
		}

		steve, err := s.User("284441212297936907")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error 404: Clark not found")
		}

		var timer *time.Timer
		var numLoopsPtr *int = &numLoops
		duration :=  10 * time.Second
		timer = time.AfterFunc(
			duration,
			func(){
				s.ChannelMessageSend(m.ChannelID, steve.Mention() + " " + arguments[1] )
				if (*numLoopsPtr != 0){
					*numLoopsPtr--
					timer.Reset(duration)
				}
			})

		s.ChannelMessageSend(m.ChannelID, "Your pingas has been installed...")
	}
}

func DdosRob(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!ddosrob") {
		commandStr := strings.TrimPrefix(m.Content, "!ddosrob ")

		const delimiter string = " "
		const numberOfArgs int = 2
		arguments := strings.SplitN(commandStr, delimiter, numberOfArgs)
		if len(arguments) != numberOfArgs {
			s.ChannelMessageSend(m.ChannelID, "ROB. WRONG FUCKING ARGUMENTS")
			return
		}

		numLoops, err := strconv.Atoi(arguments[0])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "ROB. THAT'S NOT A FUCKING INT")
			return
		}

		rob, err := s.User("293878271844679680")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error 404: Rob not found")
		}

		var timer *time.Timer
		var numLoopsPtr *int = &numLoops
		duration :=  1 * time.Millisecond
		timer = time.AfterFunc(
			duration,
			func(){
				s.ChannelMessageSend(m.ChannelID, rob.Mention() + " " + arguments[1] )
				if (*numLoopsPtr != 0){
					*numLoopsPtr--
					timer.Reset(duration)
				}
			})

		s.ChannelMessageSend(m.ChannelID, "Your pingas has been installed...")
	}
}

func KickMe(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!kickme") {
		s.GuildMemberDelete(m.GuildID, m.Author.ID)
	}
}

func MoveRob(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!moverob") {


		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 10; i++ {
			channelId := rand.Intn(4)
			switch channelId {
			case 1:
				{
					go s.GuildMemberMove(m.GuildID, rob, danish)
					go s.GuildMemberMove(m.GuildID, will, danish)
					go s.GuildMemberMove(m.GuildID, steve, danish)
				}
			case 2:
				{
					go s.GuildMemberMove(m.GuildID, rob, cowboy)
					go s.GuildMemberMove(m.GuildID, will, cowboy)
					go s.GuildMemberMove(m.GuildID, steve, cowboy)
				}
			case 3:
				{
					go s.GuildMemberMove(m.GuildID, rob, die)
					go s.GuildMemberMove(m.GuildID, will, die)
					go s.GuildMemberMove(m.GuildID, steve, die)
				}
			case 4:
				{
					go s.GuildMemberMove(m.GuildID, rob, join)
					go s.GuildMemberMove(m.GuildID, will, join)
					go s.GuildMemberMove(m.GuildID, steve, join)
				}
			}

			time.Sleep(time.Second)
		}
	}
}

func EscapeFromClarkov(s *discordgo.Session, j *discordgo.VoiceStateUpdate) {
	if clarkovEnabled {
		if j.UserID == clark {
			for i, _ := range channelArr {
				if channelArr[i] == j.ChannelID {
					if i == len(channelArr)-1 {
						i = -1
					}
					go s.GuildMemberMove(j.GuildID, steve, channelArr[i+1])
					go s.GuildMemberMove(j.GuildID, will, channelArr[i+1])
					go s.GuildMemberMove(j.GuildID, rob, channelArr[i+1])
				}
			}
		}
	}
}

func ToggleClarov(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == clark {
		clark, err := s.User(clark)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error 404: Clark not found")
		}
		s.ChannelMessageSend(m.ChannelID, "No touchy " + clark.Mention())
		return
	}
	if strings.HasPrefix(m.Content, "!toggleclarkov") {
		clarkovEnabled = !clarkovEnabled

		s.ChannelMessageSend(m.ChannelID, "EnableClarkov is " + strconv.FormatBool(clarkovEnabled))
	}
}