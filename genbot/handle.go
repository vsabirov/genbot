package genbot

// Interface that is passed to genbot in order to set custom handlers for received messages.
type MessageHandlers interface {
	OnRequestLocations(message Message)

	OnGameAnnounce(message Message)
	OnLobbyAnnounce(message Message)

	OnRequestJoin(message Message)
	OnJoinAccept(message Message)
	OnJoinDeny(message Message)

	OnRequestGameLeave(message Message)
	OnRequestLobbyLeave(message Message)

	OnSetAccept(message Message)

	OnMapAvailability(message Message)

	OnChat(message Message)

	OnGameStart(message Message)
	OnGameTimer(message Message)
	OnGameOptions(message Message)

	OnSetActive(message Message)

	OnRequestGameInfo(message Message)
}

// Default empty implementations for all message handlers.
type DefaultMessageHandlers struct{}

func (handlers DefaultMessageHandlers) OnRequestLocations(message Message) {
	announceSelf(message.Connection, message.Sender, message.BotInfo)
}

func (handlers DefaultMessageHandlers) OnGameAnnounce(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnLobbyAnnounce(message Message) {
	announceSelf(message.Connection, message.Sender, message.BotInfo)
}

func (handlers DefaultMessageHandlers) OnRequestJoin(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnJoinAccept(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnJoinDeny(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnRequestGameLeave(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnRequestLobbyLeave(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnSetAccept(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnMapAvailability(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnChat(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnGameStart(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnGameTimer(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnGameOptions(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnSetActive(message Message) {
	return
}

func (handlers DefaultMessageHandlers) OnRequestGameInfo(message Message) {
	return
}
