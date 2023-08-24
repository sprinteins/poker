<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import { PlayerType, api, stash, type Player, Status, type Set } from "$lib/api";
    import { ButtonToggle } from "$lib/button-toggle"
    import AutoReveal from "./auto-reveal.svelte"
    import CardSwitcher from "./card-switcher.svelte"
    import type { Card } from "$lib/components/cards"


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
    // Auto Reveal Timer
    //
    export let autoRevealTimer: number = 0
    function onAutoRevealTimerChange(event: Event){
        const autoRevealTimerEl = event.target as HTMLInputElement
        const autoRevealTimer = Number(autoRevealTimerEl.value) as number
        api.setAutoRevealTimer(autoRevealTimer)
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
        dispatch('done')
    }

    // 
    // Card Back Switcher
    // 
    let isCardSwitcherOpen = false
    function openCardSwitcherDialog(){
        isCardSwitcherOpen = true
    }
    function handleSwitchCardBack(cardback: Card){
        api.switchCardBack(cardback.value)
        stash.storeCardBack(cardback.value)
        isCardSwitcherOpen = false
        dispatchDone()
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

    {#if autoReveal}
        <label>
            <span class="time-box-label">Time Box:</span>
            <input type="number" on:input={onAutoRevealTimerChange} 
                value={autoRevealTimer} />
        </label>
    {/if}

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

    <button on:click={openCardSwitcherDialog}>
        Switch Card Back
    </button>

    </div>

    <div class="done-container">
        <button class="text" on:click={dispatchDone}>Done ⬆</button>
    </div>
</room-settings>

<CardSwitcher 
    open={isCardSwitcherOpen} 
    on:switchcardback={(event) => handleSwitchCardBack(event.detail)} 
/>


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

    .time-box-label{
        letter-spacing: 0.1em;
        font-family:    var(--font-all-capital);
        text-transform: uppercase;
    }
</style>