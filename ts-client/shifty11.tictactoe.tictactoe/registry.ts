import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateGame } from "./types/tictactoe/tx";
import { MsgAcceptInvite } from "./types/tictactoe/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/shifty11.tictactoe.tictactoe.MsgCreateGame", MsgCreateGame],
    ["/shifty11.tictactoe.tictactoe.MsgAcceptInvite", MsgAcceptInvite],
    
];

export { msgTypes }