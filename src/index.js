const dotenv = require("dotenv").config()
const figlet = require("figlet")
const discordjs = require("discord.js")

const client = new discordjs.Client({
	intents: ["GUILD_MESSAGES"]
})

client.on("ready", () => {
	figlet("XI_JINPING_BOT", "Swamp Land", (err, result) => {
		if(err){
			console.log(err)
		}
		console.log(result)
	})
})

client.login(process.env.DISCORD_TOKEN)