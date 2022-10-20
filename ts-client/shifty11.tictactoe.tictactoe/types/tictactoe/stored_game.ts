/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "shifty11.tictactoe.tictactoe";

export interface StoredGame {
  index: string;
  player1: string;
  player2: string;
  turn: number;
  status: StoredGame_GameStatus;
  winner: StoredGame_WinnerStatus;
  board: string;
}

export enum StoredGame_GameStatus {
  OPEN = 0,
  IN_PROGRESS = 1,
  FINISHED = 2,
  UNRECOGNIZED = -1,
}

export function storedGame_GameStatusFromJSON(
  object: any
): StoredGame_GameStatus {
  switch (object) {
    case 0:
    case "OPEN":
      return StoredGame_GameStatus.OPEN;
    case 1:
    case "IN_PROGRESS":
      return StoredGame_GameStatus.IN_PROGRESS;
    case 2:
    case "FINISHED":
      return StoredGame_GameStatus.FINISHED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return StoredGame_GameStatus.UNRECOGNIZED;
  }
}

export function storedGame_GameStatusToJSON(
  object: StoredGame_GameStatus
): string {
  switch (object) {
    case StoredGame_GameStatus.OPEN:
      return "OPEN";
    case StoredGame_GameStatus.IN_PROGRESS:
      return "IN_PROGRESS";
    case StoredGame_GameStatus.FINISHED:
      return "FINISHED";
    default:
      return "UNKNOWN";
  }
}

export enum StoredGame_WinnerStatus {
  NONE = 0,
  PLAYER1 = 1,
  PLAYER2 = 2,
  DRAW = 3,
  UNRECOGNIZED = -1,
}

export function storedGame_WinnerStatusFromJSON(
  object: any
): StoredGame_WinnerStatus {
  switch (object) {
    case 0:
    case "NONE":
      return StoredGame_WinnerStatus.NONE;
    case 1:
    case "PLAYER1":
      return StoredGame_WinnerStatus.PLAYER1;
    case 2:
    case "PLAYER2":
      return StoredGame_WinnerStatus.PLAYER2;
    case 3:
    case "DRAW":
      return StoredGame_WinnerStatus.DRAW;
    case -1:
    case "UNRECOGNIZED":
    default:
      return StoredGame_WinnerStatus.UNRECOGNIZED;
  }
}

export function storedGame_WinnerStatusToJSON(
  object: StoredGame_WinnerStatus
): string {
  switch (object) {
    case StoredGame_WinnerStatus.NONE:
      return "NONE";
    case StoredGame_WinnerStatus.PLAYER1:
      return "PLAYER1";
    case StoredGame_WinnerStatus.PLAYER2:
      return "PLAYER2";
    case StoredGame_WinnerStatus.DRAW:
      return "DRAW";
    default:
      return "UNKNOWN";
  }
}

const baseStoredGame: object = {
  index: "",
  player1: "",
  player2: "",
  turn: 0,
  status: 0,
  winner: 0,
  board: "",
};

export const StoredGame = {
  encode(message: StoredGame, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.player1 !== "") {
      writer.uint32(18).string(message.player1);
    }
    if (message.player2 !== "") {
      writer.uint32(26).string(message.player2);
    }
    if (message.turn !== 0) {
      writer.uint32(32).uint64(message.turn);
    }
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

  decode(input: Reader | Uint8Array, length?: number): StoredGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredGame } as StoredGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.player1 = reader.string();
          break;
        case 3:
          message.player2 = reader.string();
          break;
        case 4:
          message.turn = longToNumber(reader.uint64() as Long);
          break;
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

  fromJSON(object: any): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.player1 !== undefined && object.player1 !== null) {
      message.player1 = String(object.player1);
    } else {
      message.player1 = "";
    }
    if (object.player2 !== undefined && object.player2 !== null) {
      message.player2 = String(object.player2);
    } else {
      message.player2 = "";
    }
    if (object.turn !== undefined && object.turn !== null) {
      message.turn = Number(object.turn);
    } else {
      message.turn = 0;
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = storedGame_GameStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = storedGame_WinnerStatusFromJSON(object.winner);
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

  toJSON(message: StoredGame): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.player1 !== undefined && (obj.player1 = message.player1);
    message.player2 !== undefined && (obj.player2 = message.player2);
    message.turn !== undefined && (obj.turn = message.turn);
    message.status !== undefined &&
      (obj.status = storedGame_GameStatusToJSON(message.status));
    message.winner !== undefined &&
      (obj.winner = storedGame_WinnerStatusToJSON(message.winner));
    message.board !== undefined && (obj.board = message.board);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredGame>): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.player1 !== undefined && object.player1 !== null) {
      message.player1 = object.player1;
    } else {
      message.player1 = "";
    }
    if (object.player2 !== undefined && object.player2 !== null) {
      message.player2 = object.player2;
    } else {
      message.player2 = "";
    }
    if (object.turn !== undefined && object.turn !== null) {
      message.turn = object.turn;
    } else {
      message.turn = 0;
    }
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
