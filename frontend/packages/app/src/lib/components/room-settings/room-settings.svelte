<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import { PlayerType, api, type Player, Status, type Set } from "$lib/api";
    import { ButtonToggle } from "$lib/button-toggle"
    import AutoReveal from "./auto-reveal.svelte"


    // 
    // Card Sets
    // 
    export let sets: Set[] = [];
    function onCardChange(event: Event){
        const selectEl = event.target as HTMLSelectElement
        const index = selectEl.value
        const cards = sets[index].cards
        api.setCards(cards)
    }

    // 
    // Auto Reveal
    // 
    export let autoReveal: boolean = false;
    function onAutoRevealChange(event: CustomEvent){
        const autoReveal = event.detail as boolean
        api.setAutoReveal(autoReveal)
    }

    // 
    // Give Admin To
    // 
    export let status = ""
    export let players: Player[] = []
    function onAdminChange(event: Event){
        const selectEl = event.target as HTMLSelectElement
        const index = selectEl.value
        const player = players[index]
        api.setAdmin(player)
    }

    // 
    // Player Type Toggle
    // 
    export let player: Player = {
        id: "",
        name: "",
        isAdmin: false,
        chosenCard: "",
        type: PlayerType.Spectator,
    }
    const playerTypes = [
        PlayerType.Player,
        PlayerType.Spectator,
    ]
    function onActivate(event: CustomEvent<number>){
        const typeIndex = event.detail;
        const type = playerTypes[typeIndex]
        api.setPlayerType(type)
    }
    $: typeIndex = playerTypes.indexOf(player?.type)

    // 
    // Done
    // 
    const dispatch = createEventDispatcher();
    function dispatchDone(){
        console.debug('[DEBUG] ', {msg:"dispacting done"} )
        dispatch('done')
    }
</script>


<room-settings>
    <div class="settings">
    {#if player?.isAdmin }

    <select disabled={status === Status.InProgress} on:change={onCardChange}>
        {#each sets as set, si}
            <option value={si}>{set.label}: {set.cards}</option>
        {/each}
    </select>

    <AutoReveal
        value={autoReveal}
        on:activate={onAutoRevealChange}
    />

    <label>
        <span class="give-admin-label">Give admin to</span>
        <select disabled={status === Status.InProgress} on:change={onAdminChange}>
            {#each players as player, pi}
                <option value={pi}>{player.name}</option>
            {/each}
        </select>
    </label>
    {/if}

    <ButtonToggle 
        labels={playerTypes} 
        on:activate={onActivate} 
        activeIndex={typeIndex}
    />
    </div>

    <div class="done-container">
        <button class="text" on:click={dispatchDone}>Done ⬆</button>
    </div>
</room-settings>


<style>
    room-settings{
        position:         relative;
        display:          flex;
        flex-direction:   row;
        flex-wrap:        nowrap;
        background-color: var(--color-black);
    }

    .settings{
        display:          flex;
        flex-direction:   row;
        flex-wrap:        wrap;
        gap:              3rem;
    }

    .done-container{
        justify-self:    end;
        flex-grow:       1;
        display:         flex;
        justify-content: end;
    }

 button {
        min-width: 10rem;
        height:    4rem;
        padding:   0 0.5rem;

    }

    select{
        width:         200px;
        text-overflow: ellipsis;
        font-size:     1rem;
        border-width:  0.25rem;
        height:        4rem;
    }

    .give-admin-label{
        letter-spacing: 0.1em;
        font-family:    var(--font-all-capital);
        text-transform: uppercase;
    }
</style>