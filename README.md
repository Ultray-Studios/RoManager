# RoManager
Open-source web-based management panel & API designed for Roblox Game Studios

⚠️ Please note: this project is not yet ready to be used, as it is in the most early development stage it can ever be in

This project provides:
- Handling Player Data
- Use of ScyllaDB for optimized performance and incredible scalability
- Use of Redis for caching active players and handling server commands
- Management features: 
 - Managing players (active and inactive players)
 - Database cleanup
 - Scheduled game updates
 - Data Backups
 - Moderation features
 - Multiple Staff Accounts
 - Database Editor
 - Monitoring
 - Logging
 - Handling GDPR requests automatically
 - Cheat detection using data analysis
- Integration into your own systems

## Project structure:
### Source code:
- src/admin: Management panel
- src/api: API that recieves and sends information to from and to Roblox gameservers 
- src/game: Lua source code that runs in Roblox gameservers: sends player data, recieves commands to execute from the management panel
