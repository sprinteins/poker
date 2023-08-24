<script lang="ts">
    import { goto } from '$app/navigation';
    import { page } from "$app/stores";
    
    
    const { 
        roomid: roomID,
    } = $page.params

    let playerName = "";
    let roomName = randomRoomName();
    
    if( roomID ){
        roomName = roomID
    }
    
    function routeToPage(route: string, replaceState: boolean = true) {
       goto(`/${route}`, { replaceState }) 
    }
    
    
    function onSubmit(event: Event){
        event.preventDefault();
    
        const path = ["room", roomName, playerName].join("/")
     
        routeToPage(path)
    }

    function randomRoomName(): string {
        const name = Math.random().toString(36).slice(2,10).toUpperCase()
        return name
    }
    
</script>


<svelte:head>
	<title>Poker: Join Room</title>
</svelte:head>

<room>
    <logo>
        <img src="/img/logo_white.svg" alt="logo" />
    </logo>
    <main>

        <form on:submit={onSubmit}>
    
            <div>
                <label for="room-name">Room Name</label>
                <input 
                    id="room-name" 
                    test-id="room-name"
                    bind:value={roomName} 
                    type="text"
                    placeholder="PC2EY"

                    
                />
            </div>

            <div>
                <label for="player-name" >Player Name</label>
                <!-- svelte-ignore a11y-autofocus -->
                <input 
                    id="player-name" 
                    test-id="player-name"
                    bind:value={playerName} 
                    type="text" 
                    placeholder="Jane" 
                    autofocus
                />
            </div>
    
    
            <div class="button-container">
                <button type="submit" test-id="button-join" class="primary">Enter</button>
            </div>
        </form>
    </main>

    <footer>

        <span class="made-by">Made without 🍪</span>
        <a href="https://www.sprinteins.com" target="_blank">
            <img src="/img/se_logo_white.png" class="se-logo" alt="SprintEins Logo"/>
        </a> 
    </footer>

</room>
    
<style>


    room {
        height:             100%;
        display:            grid;
        grid-template-rows: auto 1fr auto;
    }

    main {
        display: flex;
        justify-content: center;
    }
    logo{
        display:    block;
        text-align: center
    }
    logo img {
        width:           900px;
        height:          220px;
        object-fit:      cover;
        object-position: 50% -150px;
        margin-top:      2.5rem;
        margin-bottom:   2.5rem;
    }
    
    form {
        display:         flex;
        gap:             5rem;
        flex-direction:  column;
        flex-wrap:       wrap;
        justify-content: center;

        width: auto;
    }

    label {
        display:        block;
        font-family:    "SoinSansProHeadline";
        letter-spacing: 0.1em;
        text-transform: uppercase;
        font-weight:    400;
        color:          white;
        font-size:      1.25rem;
        margin-bottom:  0.5rem;
    }

    input {
        /* height:    120px; */
        width:        300px;
        font-size:    1.25rem;
        padding:      1.25rem;
        border-width: 0.25rem;
    }

    .button-container {
        text-align: right;
    }

    button {
        padding:      0 1.5rem;
        border-width: 0.25rem;
    }

    footer {
		display: 		 flex;
		flex-direction:  column;
		justify-content: center;
		align-items: 	 center;
		padding: 		 40px;
	}

	.made-by {
		color: 	   		white;
		font-size: 		1rem;
		font-family: 	var(--font-all-capital);
		text-transform: uppercase;
        letter-spacing: 0.1em;
    }

	footer a {
		font-weight: bold;
	}

	.se-logo{
		width: 200px;
	}


</style>