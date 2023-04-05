# Overview
This is a LAN lobby chat bot for C&C Generals: Zero Hour. It reads UDP packets from the LAN interface specified by an address, parses them and reacts appropriately.
You need a LAN (emulated or not, e.g. Radmin, Hamachi) with active players, and a different machine with the game installed available to test this bot.

# Starting
Connect both the bot machine and your game machine to the same LAN network. Start the bot, tell it the address of the LAN interface with port 8086 (default ZH port for LAN lobby) when prompted.

Start the game, open the multiplayer network window. A player named 'Genbot' should be visible in the player list if everything worked correctly.

# Using
Current commands include:

- !guess - Used in main lobby chat. Guess a number from 0 to 100 with binary search hints.
- !announce - Used in private game lobbies. You can announce any message you want from your game lobby into the main LAN lobby for all players to see.
