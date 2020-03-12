package delegates

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func KieranSpeak(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!story" {
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