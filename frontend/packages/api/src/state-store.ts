import { writable } from "svelte/store";

export enum Status {
    Init       = "Init",     
    Start      = "Start",
    InProgress = "InProgress",
    Revealed   = "Revealed",
}

export type Cards = string[]
export type Set = {
    label: string
    cards: Cards
}

export type State = {
    player?:            Player
    players?:           Player[]
    status?:            Status
    cards:              Cards
    sets:               Set[]
    autoReveal:         boolean
    autoRevealTimeout:  number
}

export interface Player {
    id:         string
    name:       string
    isAdmin:    boolean
    chosenCard: string
    type:       PlayerType
    cardBack:   string
}

export enum PlayerType {
    Spectator = "Spectator",
    Player    = "Player", 
}

const initalState: State = {
    players:           [],
    status:            Status.Start,
    cards:             [],
    sets:              [],
    autoReveal:        false,
    autoRevealTimeout: 0,
}

const state = writable(initalState)

export type UnsubscriberFn = () => void
export type HandlerFn = (state: State) => void 
export type HandlerFnOnOpen = () => void

export function subscribe(handler: HandlerFn): UnsubscriberFn {
    return state.subscribe(handler)
}

export const setState = state.set