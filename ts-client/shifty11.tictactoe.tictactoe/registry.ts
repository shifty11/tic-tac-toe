import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAcceptInvite } from "./types/tictactoe/tx";
import { MsgCreateGame } from "./types/tictactoe/tx";
import { MsgPlayTurn } from "./types/tictactoe/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/shifty11.tictactoe.tictactoe.MsgAcceptInvite", MsgAcceptInvite],
    ["/shifty11.tictactoe.tictactoe.MsgCreateGame", MsgCreateGame],
    ["/shifty11.tictactoe.tictactoe.MsgPlayTurn", MsgPlayTurn],
    
];

export { msgTypes }