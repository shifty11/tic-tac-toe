/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "shifty11.tictactoe.tictactoe";

export interface EventCreateGame {
  gameIndex: number;
  player1: string;
  player2: string;
}

export interface EventInviteAccepted {
  gameIndex: number;
  player1: string;
  player2: string;
}

export interface EventTurnPlayed {
  gameIndex: number;
  playedBy: string;
  row: number;
  column: number;
  winner: EventTurnPlayed_WinnerStatus;
  board: string;
}

export enum EventTurnPlayed_WinnerStatus {
  NONE = 0,
  PLAYER1 = 1,
  PLAYER2 = 2,
  DRAW = 3,
  UNRECOGNIZED = -1,
}

export function eventTurnPlayed_WinnerStatusFromJSON(
  object: any
): EventTurnPlayed_WinnerStatus {
  switch (object) {
    case 0:
    case "NONE":
      return EventTurnPlayed_WinnerStatus.NONE;
    case 1:
    case "PLAYER1":
      return EventTurnPlayed_WinnerStatus.PLAYER1;
    case 2:
    case "PLAYER2":
      return EventTurnPlayed_WinnerStatus.PLAYER2;
    case 3:
    case "DRAW":
      return EventTurnPlayed_WinnerStatus.DRAW;
    case -1:
    case "UNRECOGNIZED":
    default:
      return EventTurnPlayed_WinnerStatus.UNRECOGNIZED;
  }
}

export function eventTurnPlayed_WinnerStatusToJSON(
  object: EventTurnPlayed_WinnerStatus
): string {
  switch (object) {
    case EventTurnPlayed_WinnerStatus.NONE:
      return "NONE";
    case EventTurnPlayed_WinnerStatus.PLAYER1:
      return "PLAYER1";
    case EventTurnPlayed_WinnerStatus.PLAYER2:
      return "PLAYER2";
    case EventTurnPlayed_WinnerStatus.DRAW:
      return "DRAW";
    default:
      return "UNKNOWN";
  }
}

const baseEventCreateGame: object = { gameIndex: 0, player1: "", player2: "" };

export const EventCreateGame = {
  encode(message: EventCreateGame, writer: Writer = Writer.create()): Writer {
    if (message.gameIndex !== 0) {
      writer.uint32(8).uint64(message.gameIndex);
    }
    if (message.player1 !== "") {
      writer.uint32(18).string(message.player1);
    }
    if (message.player2 !== "") {
      writer.uint32(26).string(message.player2);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventCreateGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventCreateGame } as EventCreateGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.player1 = reader.string();
          break;
        case 3:
          message.player2 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventCreateGame {
    const message = { ...baseEventCreateGame } as EventCreateGame;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = Number(object.gameIndex);
    } else {
      message.gameIndex = 0;
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
    return message;
  },

  toJSON(message: EventCreateGame): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.player1 !== undefined && (obj.player1 = message.player1);
    message.player2 !== undefined && (obj.player2 = message.player2);
    return obj;
  },

  fromPartial(object: DeepPartial<EventCreateGame>): EventCreateGame {
    const message = { ...baseEventCreateGame } as EventCreateGame;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = 0;
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
    return message;
  },
};

const baseEventInviteAccepted: object = {
  gameIndex: 0,
  player1: "",
  player2: "",
};

export const EventInviteAccepted = {
  encode(
    message: EventInviteAccepted,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gameIndex !== 0) {
      writer.uint32(8).uint64(message.gameIndex);
    }
    if (message.player1 !== "") {
      writer.uint32(18).string(message.player1);
    }
    if (message.player2 !== "") {
      writer.uint32(26).string(message.player2);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventInviteAccepted {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventInviteAccepted } as EventInviteAccepted;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.player1 = reader.string();
          break;
        case 3:
          message.player2 = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventInviteAccepted {
    const message = { ...baseEventInviteAccepted } as EventInviteAccepted;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = Number(object.gameIndex);
    } else {
      message.gameIndex = 0;
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
    return message;
  },

  toJSON(message: EventInviteAccepted): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.player1 !== undefined && (obj.player1 = message.player1);
    message.player2 !== undefined && (obj.player2 = message.player2);
    return obj;
  },

  fromPartial(object: DeepPartial<EventInviteAccepted>): EventInviteAccepted {
    const message = { ...baseEventInviteAccepted } as EventInviteAccepted;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = 0;
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
    return message;
  },
};

const baseEventTurnPlayed: object = {
  gameIndex: 0,
  playedBy: "",
  row: 0,
  column: 0,
  winner: 0,
  board: "",
};

export const EventTurnPlayed = {
  encode(message: EventTurnPlayed, writer: Writer = Writer.create()): Writer {
    if (message.gameIndex !== 0) {
      writer.uint32(8).uint64(message.gameIndex);
    }
    if (message.playedBy !== "") {
      writer.uint32(18).string(message.playedBy);
    }
    if (message.row !== 0) {
      writer.uint32(24).uint32(message.row);
    }
    if (message.column !== 0) {
      writer.uint32(32).uint32(message.column);
    }
    if (message.winner !== 0) {
      writer.uint32(40).int32(message.winner);
    }
    if (message.board !== "") {
      writer.uint32(50).string(message.board);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventTurnPlayed {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventTurnPlayed } as EventTurnPlayed;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameIndex = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.playedBy = reader.string();
          break;
        case 3:
          message.row = reader.uint32();
          break;
        case 4:
          message.column = reader.uint32();
          break;
        case 5:
          message.winner = reader.int32() as any;
          break;
        case 6:
          message.board = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventTurnPlayed {
    const message = { ...baseEventTurnPlayed } as EventTurnPlayed;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = Number(object.gameIndex);
    } else {
      message.gameIndex = 0;
    }
    if (object.playedBy !== undefined && object.playedBy !== null) {
      message.playedBy = String(object.playedBy);
    } else {
      message.playedBy = "";
    }
    if (object.row !== undefined && object.row !== null) {
      message.row = Number(object.row);
    } else {
      message.row = 0;
    }
    if (object.column !== undefined && object.column !== null) {
      message.column = Number(object.column);
    } else {
      message.column = 0;
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = eventTurnPlayed_WinnerStatusFromJSON(object.winner);
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

  toJSON(message: EventTurnPlayed): unknown {
    const obj: any = {};
    message.gameIndex !== undefined && (obj.gameIndex = message.gameIndex);
    message.playedBy !== undefined && (obj.playedBy = message.playedBy);
    message.row !== undefined && (obj.row = message.row);
    message.column !== undefined && (obj.column = message.column);
    message.winner !== undefined &&
      (obj.winner = eventTurnPlayed_WinnerStatusToJSON(message.winner));
    message.board !== undefined && (obj.board = message.board);
    return obj;
  },

  fromPartial(object: DeepPartial<EventTurnPlayed>): EventTurnPlayed {
    const message = { ...baseEventTurnPlayed } as EventTurnPlayed;
    if (object.gameIndex !== undefined && object.gameIndex !== null) {
      message.gameIndex = object.gameIndex;
    } else {
      message.gameIndex = 0;
    }
    if (object.playedBy !== undefined && object.playedBy !== null) {
      message.playedBy = object.playedBy;
    } else {
      message.playedBy = "";
    }
    if (object.row !== undefined && object.row !== null) {
      message.row = object.row;
    } else {
      message.row = 0;
    }
    if (object.column !== undefined && object.column !== null) {
      message.column = object.column;
    } else {
      message.column = 0;
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
