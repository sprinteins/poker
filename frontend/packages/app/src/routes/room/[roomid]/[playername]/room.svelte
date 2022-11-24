<script lang="ts">
	import { page } from "$app/stores"; 
	import { api, type Cards as TypeCards, type Set, Status, type Player, PlayerType } from "$lib/api"
	import { Cards } from "$lib/components/cards"
	import { PlayerList } from "$lib/components/player-list"
	import { RoomController } from "$lib/components/room-controller";
	
	const { 
		roomid:     roomID,
		playername: playerName,
	} = $page.params
	
	let player: Player | undefined;
	let players: Player[] | undefined = []
	let roomStatus: Status | undefined = Status.Init
	let cards: TypeCards = []
	let sets: Set[] = []
	let autoReveal: boolean = false
	
	api.register(roomID, playerName, (data) => {
		
		player = data.player
		players = data.players
		roomStatus = data.status
		cards = data.cards
		sets = data.sets
		autoReveal = data.autoReveal
		
	})
	
	function onChoose(card: string){
		api.choose(card)
	}
	
	function onRevote(){
		if(!isVotingInProgress(roomStatus)){
			return
		}
		api.reVote()
	}
	
	function hasPlayerChosen(player: Player): boolean {
		if(!player){
			return false
		}
		
		return player.chosenCard !== "" 
	}
	
	function isGameRunnin(roomStatus: Status): boolean {
		return roomStatus === Status.InProgress
	}
	
	function isVotingInProgress(roomStatus: Status): boolean {
		return roomStatus === Status.InProgress
	}
	
	function isRevealed(roomStatus: Status): boolean {
		return roomStatus === Status.Revealed
	}
	
	function isPlayerSpectator( player: Player): boolean{
		return player && player.type === PlayerType.Spectator
	}
	
</script>

<room>
	
	<RoomController
	status={roomStatus} 
	player={player} 
	players={players}
	roomID={roomID} 
	sets={sets} 
	autoReveal={autoReveal}    
	/>
	
	
	
	{#if isPlayerSpectator(player) || hasPlayerChosen(player) || !isVotingInProgress(roomStatus)}
	<PlayerList 
	players={players} 
	currentPlayer={player} 
	isRevealed={isRevealed(roomStatus)}
	on:revote={onRevote}
	/>
	{/if}
	
	{#if !isPlayerSpectator(player) && ( !hasPlayerChosen(player) && isGameRunnin(roomStatus) )}
	<Cards 
	cards={cards} 
	on:choose={(event) => onChoose(event.detail)} 
	/>
	{/if}
	
	<footer>
		<span> &nbsp; </span>
		<a href="https://www.sprinteins.com" class="logo" target="_blank">
			<img src="/img/logo_white.svg" class="logo" alt="Scrum-Poker Logo"/>
		</a> 
		
		<a href="https://www.sprinteins.com" target="_blank" class="se-logo">
			<img src="/img/se_logo_white.png" alt="SprintEins Logo"/>
		</a> 
	</footer>
	
	
</room>

<style>
	room {
		height:             100%;
		display:            grid;
		gap:                2rem;
		grid-template-rows: auto 1fr auto;
	}
	
	footer {
		display:     grid;
		align-items: baseline;
		height:      190px;
		padding:     0;
		grid-template-columns: 1fr 1fr 200px;
	}
	
	
	a.logo{
		width:      200px;
		height:     inherit;
		display:    block;
		text-align: center;
		padding:    0;
	}
	
	a.logo img {
		width:           200px;
		height:          inherit;
		object-fit:      none;
		object-position: 50% 71%;
	}
	
	.se-logo{
		width: 200px;
	}
	
	.se-logo img{
		width: inherit;
	}
	
	
</style>