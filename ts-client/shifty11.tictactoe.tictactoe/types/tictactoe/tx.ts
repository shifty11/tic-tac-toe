/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "shifty11.tictactoe.tictactoe";

export interface MsgCreateGame {
  creator: string;
  player2: string;
}

export interface MsgCreateGameResponse {
  gameId: number;
}

export interface MsgAcceptInvite {
  creator: string;
  gameId: number;
}

export interface MsgAcceptInviteResponse {}

export interface MsgPlayTurn {
  creator: string;
  gameId: number;
  x: number;
  y: number;
}

export interface MsgPlayTurnResponse {
  status: MsgPlayTurnResponse_GameStatus;
  winner: MsgPlayTurnResponse_WinnerStatus;
  board: string;
}

export enum MsgPlayTurnResponse_GameStatus {
  OPEN = 0,
  IN_PROGRESS = 1,
  FINISHED = 2,
  UNRECOGNIZED = -1,
}

export function msgPlayTurnResponse_GameStatusFromJSON(
  object: any
): MsgPlayTurnResponse_GameStatus {
  switch (object) {
    case 0:
    case "OPEN":
      return MsgPlayTurnResponse_GameStatus.OPEN;
    case 1:
    case "IN_PROGRESS":
      return MsgPlayTurnResponse_GameStatus.IN_PROGRESS;
    case 2:
    case "FINISHED":
      return MsgPlayTurnResponse_GameStatus.FINISHED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return MsgPlayTurnResponse_GameStatus.UNRECOGNIZED;
  }
}

export function msgPlayTurnResponse_GameStatusToJSON(
  object: MsgPlayTurnResponse_GameStatus
): string {
  switch (object) {
    case MsgPlayTurnResponse_GameStatus.OPEN:
      return "OPEN";
    case MsgPlayTurnResponse_GameStatus.IN_PROGRESS:
      return "IN_PROGRESS";
    case MsgPlayTurnResponse_GameStatus.FINISHED:
      return "FINISHED";
    default:
      return "UNKNOWN";
  }
}

export enum MsgPlayTurnResponse_WinnerStatus {
  NONE = 0,
  PLAYER1 = 1,
  PLAYER2 = 2,
  DRAW = 3,
  UNRECOGNIZED = -1,
}

export function msgPlayTurnResponse_WinnerStatusFromJSON(
  object: any
): MsgPlayTurnResponse_WinnerStatus {
  switch (object) {
    case 0:
    case "NONE":
      return MsgPlayTurnResponse_WinnerStatus.NONE;
    case 1:
    case "PLAYER1":
      return MsgPlayTurnResponse_WinnerStatus.PLAYER1;
    case 2:
    case "PLAYER2":
      return MsgPlayTurnResponse_WinnerStatus.PLAYER2;
    case 3:
    case "DRAW":
      return MsgPlayTurnResponse_WinnerStatus.DRAW;
    case -1:
    case "UNRECOGNIZED":
    default:
      return MsgPlayTurnResponse_WinnerStatus.UNRECOGNIZED;
  }
}

export function msgPlayTurnResponse_WinnerStatusToJSON(
  object: MsgPlayTurnResponse_WinnerStatus
): string {
  switch (object) {
    case MsgPlayTurnResponse_WinnerStatus.NONE:
      return "NONE";
    case MsgPlayTurnResponse_WinnerStatus.PLAYER1:
      return "PLAYER1";
    case MsgPlayTurnResponse_WinnerStatus.PLAYER2:
      return "PLAYER2";
    case MsgPlayTurnResponse_WinnerStatus.DRAW:
      return "DRAW";
    default:
      return "UNKNOWN";
  }
}

const baseMsgCreateGame: object = { creator: "", player2: "" };

export const MsgCreateGame = {
  encode(message: MsgCreateGame, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.player2 !== "") {
      writer.uint32(18).string(message.player2);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateGame } as MsgCreateGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.player2 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGame {
    const message = { ...baseMsgCreateGame } as MsgCreateGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.player2 !== undefined && object.player2 !== null) {
      message.player2 = String(object.player2);
    } else {
      message.player2 = "";
    }
    return message;
  },

  toJSON(message: MsgCreateGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.player2 !== undefined && (obj.player2 = message.player2);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateGame>): MsgCreateGame {
    const message = { ...baseMsgCreateGame } as MsgCreateGame;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.player2 !== undefined && object.player2 !== null) {
      message.player2 = object.player2;
    } else {
      message.player2 = "";
    }
    return message;
  },
};

const baseMsgCreateGameResponse: object = { gameId: 0 };

export const MsgCreateGameResponse = {
  encode(
    message: MsgCreateGameResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameId !== 0) {
      writer.uint32(8).uint64(message.gameId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateGameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGameResponse {
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse;
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId);
    } else {
      message.gameId = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateGameResponse): unknown {
    const obj: any = {};
    message.gameId !== undefined && (obj.gameId = message.gameId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateGameResponse>
  ): MsgCreateGameResponse {
    const message = { ...baseMsgCreateGameResponse } as MsgCreateGameResponse;
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId;
    } else {
      message.gameId = 0;
    }
    return message;
  },
};

const baseMsgAcceptInvite: object = { creator: "", gameId: 0 };

export const MsgAcceptInvite = {
  encode(message: MsgAcceptInvite, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameId !== 0) {
      writer.uint32(16).uint64(message.gameId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAcceptInvite {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAcceptInvite } as MsgAcceptInvite;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.gameId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAcceptInvite {
    const message = { ...baseMsgAcceptInvite } as MsgAcceptInvite;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId);
    } else {
      message.gameId = 0;
    }
    return message;
  },

  toJSON(message: MsgAcceptInvite): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameId !== undefined && (obj.gameId = message.gameId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAcceptInvite>): MsgAcceptInvite {
    const message = { ...baseMsgAcceptInvite } as MsgAcceptInvite;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId;
    } else {
      message.gameId = 0;
    }
    return message;
  },
};

const baseMsgAcceptInviteResponse: object = {};

export const MsgAcceptInviteResponse = {
  encode(_: MsgAcceptInviteResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAcceptInviteResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAcceptInviteResponse,
    } as MsgAcceptInviteResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAcceptInviteResponse {
    const message = {
      ...baseMsgAcceptInviteResponse,
    } as MsgAcceptInviteResponse;
    return message;
  },

  toJSON(_: MsgAcceptInviteResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAcceptInviteResponse>
  ): MsgAcceptInviteResponse {
    const message = {
      ...baseMsgAcceptInviteResponse,
    } as MsgAcceptInviteResponse;
    return message;
  },
};

const baseMsgPlayTurn: object = { creator: "", gameId: 0, x: 0, y: 0 };

export const MsgPlayTurn = {
  encode(message: MsgPlayTurn, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameId !== 0) {
      writer.uint32(16).uint64(message.gameId);
    }
    if (message.x !== 0) {
      writer.uint32(24).uint64(message.x);
    }
    if (message.y !== 0) {
      writer.uint32(32).uint64(message.y);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPlayTurn {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPlayTurn } as MsgPlayTurn;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.gameId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.x = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.y = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlayTurn {
    const message = { ...baseMsgPlayTurn } as MsgPlayTurn;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = Number(object.gameId);
    } else {
      message.gameId = 0;
    }
    if (object.x !== undefined && object.x !== null) {
      message.x = Number(object.x);
    } else {
      message.x = 0;
    }
    if (object.y !== undefined && object.y !== null) {
      message.y = Number(object.y);
    } else {
      message.y = 0;
    }
    return message;
  },

  toJSON(message: MsgPlayTurn): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameId !== undefined && (obj.gameId = message.gameId);
    message.x !== undefined && (obj.x = message.x);
    message.y !== undefined && (obj.y = message.y);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPlayTurn>): MsgPlayTurn {
    const message = { ...baseMsgPlayTurn } as MsgPlayTurn;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.gameId !== undefined && object.gameId !== null) {
      message.gameId = object.gameId;
    } else {
      message.gameId = 0;
    }
    if (object.x !== undefined && object.x !== null) {
      message.x = object.x;
    } else {
      message.x = 0;
    }
    if (object.y !== undefined && object.y !== null) {
      message.y = object.y;
    } else {
      message.y = 0;
    }
    return message;
  },
};

const baseMsgPlayTurnResponse: object = { status: 0, winner: 0, board: "" };

export const MsgPlayTurnResponse = {
  encode(
    message: MsgPlayTurnResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.status !== 0) {
      writer.uint32(40).int32(message.status);
    }
    if (message.winner !== 0) {
      writer.uint32(48).int32(message.winner);
    }
    if (message.board !== "") {
      writer.uint32(58).string(message.board);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPlayTurnResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPlayTurnResponse } as MsgPlayTurnResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 5:
          message.status = reader.int32() as any;
          break;
        case 6:
          message.winner = reader.int32() as any;
          break;
        case 7:
          message.board = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlayTurnResponse {
    const message = { ...baseMsgPlayTurnResponse } as MsgPlayTurnResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = msgPlayTurnResponse_GameStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = msgPlayTurnResponse_WinnerStatusFromJSON(object.winner);
    } else {
      message.winner = 0;
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = String(object.board);
    } else {
      message.board = "";
    }
    return message;
  },

  toJSON(message: MsgPlayTurnResponse): unknown {
    const obj: any = {};
    message.status !== undefined &&
      (obj.status = msgPlayTurnResponse_GameStatusToJSON(message.status));
    message.winner !== undefined &&
      (obj.winner = msgPlayTurnResponse_WinnerStatusToJSON(message.winner));
    message.board !== undefined && (obj.board = message.board);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPlayTurnResponse>): MsgPlayTurnResponse {
    const message = { ...baseMsgPlayTurnResponse } as MsgPlayTurnResponse;
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = 0;
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = object.board;
    } else {
      message.board = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
  AcceptInvite(request: MsgAcceptInvite): Promise<MsgAcceptInviteResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  PlayTurn(request: MsgPlayTurn): Promise<MsgPlayTurnResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse> {
    const data = MsgCreateGame.encode(request).finish();
    const promise = this.rpc.request(
      "shifty11.tictactoe.tictactoe.Msg",
      "CreateGame",
      data
    );
    return promise.then((data) =>
      MsgCreateGameResponse.decode(new Reader(data))
    );
  }

  AcceptInvite(request: MsgAcceptInvite): Promise<MsgAcceptInviteResponse> {
    const data = MsgAcceptInvite.encode(request).finish();
    const promise = this.rpc.request(
      "shifty11.tictactoe.tictactoe.Msg",
      "AcceptInvite",
      data
    );
    return promise.then((data) =>
      MsgAcceptInviteResponse.decode(new Reader(data))
    );
  }

  PlayTurn(request: MsgPlayTurn): Promise<MsgPlayTurnResponse> {
    const data = MsgPlayTurn.encode(request).finish();
    const promise = this.rpc.request(
      "shifty11.tictactoe.tictactoe.Msg",
      "PlayTurn",
      data
    );
    return promise.then((data) => MsgPlayTurnResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
