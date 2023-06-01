import type { GameState } from "./constants";

export type User = {
	email: string;
	username: string;
	verified: boolean;
	token: string;
};

export type CreateGameResponse = {
	id: string;
	statusCode: number;
};

export type Player = {
	avatar: string;
	id: string;
	username: string;
	score: number;
	picking: boolean;
	bids: number;
};

export type DealResponse = {
	round: number;
	trick: number;
	cards: number[];
	state: GameState;
}