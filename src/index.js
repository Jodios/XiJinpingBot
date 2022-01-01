const dotenv = require("dotenv").config()
const figlet = require("figlet")
const discordjs = require("discord.js")
const firewall = require("./resources/firewall.json")

const client = new discordjs.Client({
	intents: ["GUILD_MESSAGES", "GUILDS"]
})

client.on("messageCreate", (message) => {
	firewall.forEach((badWord, i) => {
		let matches = message.content.toLowerCase().match(`.*${badWord.toLowerCase()}.*`)
		if(matches != undefined && matches.length != 0 ){
			message.channel.send(":angry:")
			return;
		}
	})

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