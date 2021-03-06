import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { GameState } from "../../../games/poker/game/state.js";
import { Game } from "../../../games/poker/connection/socket.js";

export const usePokerTicketRequest = async (args) => {
    const params = new URLSearchParams({ buyIn: args.buyIn, class: args.class });
    const res = await _fetch(_getUrl(args.id), params);
    if (res.ok) {
        return await res.json();
    } else {
        return Promise.reject();
    }
    //return Promise.resolve();
};

const _getUrl = (id) => {
    let url = "";
    if (process.env.NEXT_PUBLIC_POKER_TICKET_HOST != undefined) {
        url = process.env.NEXT_PUBLIC_POKER_TICKET_HOST;
    }
    if (id) {
        console.log("Requesting ticket with ID");
        return `${url}/api/poker/ticket/${id}`;
    }
    console.log("Requesting ticket without ID");
    return `${url}/api/poker/ticket`;
};

const _fetch = async (url, params) => {
    return fetch(`${url}?${params.toString()}`, {
        mode: "cors",
        credentials: "include",
        method: "GET"
    });
};

export const usePokerMatch = () => {
    const [game, setGame] = useState({});
    const [joined, setJoined] = useState(false);
    const [actions, setActions] = useState({});
    const router = useRouter();
    const { lobbyId, buyInClass, buyIn } = router.query;

    useEffect(async () => {
        const params = new URLSearchParams({ buyIn: buyIn, class: buyInClass });
        const res = await _fetch(_getUrl(lobbyId), params);

        if (!res.ok) {
            await router.push("/games/poker");
            return;
        }
        const ticket = await res.json();
        console.log("Ticket: ", ticket);
        if (!ticket.address || !ticket.token) {
            await router.push("/games/poker");
            return;
        }
        let gameState = new GameState();
        gameState.setOnPossibleActions((actions) => {
            setActions(actions);
        });
        let game = new Game(gameState, ticket, () => {
            router.push("/games/poker").then();
        });
        game.start();
        console.log("Starting");
        setGame(game);
        setJoined(true);

        return () => {
            console.log("Closing Websocket connection");
            game.close(1001, "");
        };
    }, []);

    return { game, joined, actions };
};
