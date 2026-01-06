import express, { Request, Response, Express } from "express";
import { Client } from "discord.js";
import { config } from "@utils/config";
import { IncomingMessage, Server, ServerResponse } from "http";

export enum HookMethods {
  Get,
  Post,
}

export default class Hooks {
  private express: Express;
  private discord: Client;

  constructor(client: Client) {
    this.express = express();
    this.express.use(express.json());

    this.discord = client;

    return this;
  }

  public addHook(
    method: HookMethods,
    path: string,
    func: (
      req: Request,
      res: Response,
      discord: Client,
    ) => void | Promise<void>,
  ): void {
    switch (method) {
      case HookMethods.Get:
        this.express.get(path, (req: Request, res: Response) =>
          func(req, res, this.discord),
        );
        break;
      case HookMethods.Post:
        this.express.post(path, async (req: Request, res: Response) =>
          func(req, res, this.discord),
        );
        break;
    }
  }

  public serve(): Server<typeof IncomingMessage, typeof ServerResponse> {
    return this.express.listen(config.WEBHOOK_PORT, () => {});
  }
}
