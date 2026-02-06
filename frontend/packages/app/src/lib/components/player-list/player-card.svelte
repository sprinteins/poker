<script lang="ts">
    import { CardBacks } from "$lib/card-backs";
    import type { Player } from "../../api";
    

    export let player: Player;
    export let isCurrentPlayer: boolean;
    export let isRevealed: boolean = false;
    $: playerHasChosen = player.chosenCard !== ""

    const now = new Date()
    const currentMonth = now.getMonth() + 1 
    const currentDay = now.getDate()

    const monthDecember = 12
    const dayNikolaus = 6

    const defaultCardBack = CardBacks.logo
    let cardBack = defaultCardBack
    $: if(player.cardBack!==""){
        cardBack = player.cardBack
    }

    $: isEasterEggSanta = currentMonth === monthDecember && currentDay === dayNikolaus
    $: isEasterEggDecember = currentMonth === monthDecember

</script>

<div 
    class="player-card" 
    class:playerHasChosen
    on:click
>
    <div class="admin">
        {#if player.isAdmin}
            {#if isEasterEggSanta}
                🎅
            {:else if isEasterEggDecember}
                🎄
            {:else}
                👑
            {/if}
        {/if}
        {#if !player.isAdmin}
            {#if isEasterEggSanta}
                <span class="reindeer">🦌</span>
            {:else if isEasterEggDecember}
                🎁
            {:else}
                &nbsp;    
            {/if}
        {/if}
    </div>    
    <div class="player" test-id="player-card__player-name" title={player.name}>
        {player.name}
    </div>
    <div 
        class="card flip-card" 
        class:revealed={isRevealed} 
    >
        <div class="flip-card-inner">
            <div 
                class="flip-card-front" 
                test-id="player-card__front"
                class:isCurrentPlayer 
            >
                <img 
                    src={cardBack} 
                    class="card-image"  
                    class:default={cardBack===defaultCardBack}
                    alt="SprintEins Logo small"
                />
            </div>
            <div 
                class="flip-card-back"  
                test-id="player-card__back"
                class:isCurrentPlayer
            >
                <div class="content">
                    {player.chosenCard}
                </div>
            </div>
        </div>
    </div>
</div>

<style>
    .isCurrentPlayer{
        --shadow: inset 0px 0px 3px 1px var(--color-red);
        -webkit-box-shadow: var(--shadow);
	    box-shadow: 		var(--shadow);
    }
    .player-card {
        font-family: "Helvetica Neue";
        transition:   all 0.3s ease-in-out;
    }
     .card {
        width:  12rem;
        height: 20rem;

        color:         white;
        cursor:        pointer; 
        transition:    all 0.3s ease;
    }

    .player-card:not(.playerHasChosen){
        opacity: 50%;
        transform: scale(0.9, 0.9);
    }

    .content {
        font-size:      clamp(1.5rem, 4vw, 3.5rem);
        font-weight:    bold;
        position:       relative;
        z-index:        1;
        width:          100%;
        height:         100%;
        display:        grid;
        place-items:    center;
        text-transform: uppercase;
        text-align:     center;
        padding:        0 0.75rem;
        line-height:    1;
        overflow-wrap:  anywhere;
        
    }

    .player{
        width:         12rem;
        text-align:    center;
        font-size:     2rem;
        color:         white;
        white-space:   nowrap;
        text-overflow: ellipsis;
        overflow:      hidden;
    }
    .admin{
        width:      12rem;
        text-align: center;
        font-size:  3rem;
        height:     3rem;
    }

/* ------- */


    .flip-card {
        border-radius:    4px;
        background-color: transparent;
        width:            12rem;
        height:           20rem;
        perspective:      1000px;
        display:          grid;
        place-items:      center;
    }

    .flip-card-inner {
        position:        relative;
        width:           100%;
        height:          100%;
        text-align:      center;
        transition:      transform 0.6s;
        transform-style: preserve-3d;
    }

    .flip-card.revealed .flip-card-inner{
        transform: rotateY(180deg);
    }

    .flip-card-front, .flip-card-back {
        position:      absolute;
        width:         100%;
        height:        100%;
        border-radius: 4px;
        display:       grid;
        place-items:   center;

        -webkit-backface-visibility: hidden;
        backface-visibility:         hidden;
    }

    .card-image {
        width: 99%;
    }
    .card-image.default{
        width: 5rem;
    }

    .flip-card-front {
        /* background-color: #bbb; */
        /* background-image: "url(./card-bg.png)"; */
        /* background-color: black; */
        background-color: var(--color-gray-dark);
        color: black;
    }

    .flip-card-back {
        /* background-color: #2980b9; */
        background-color: var(--color-gray-dark);
        color: white;
        transform: rotateY(180deg);
    }


    /* EASTER EGGS */

    .reindeer {
        transform: rotateY(180deg);
        display:   inline-block;
    }

</style>