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

/** Msg defines the Msg service. */
export interface Msg {
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AcceptInvite(request: MsgAcceptInvite): Promise<MsgAcceptInviteResponse>;
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
