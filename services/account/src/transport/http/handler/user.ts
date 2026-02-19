import { Context } from "hono";
import { UserRepository } from "../../../ports/repo";

export const get = (repo: UserRepository) => {
    return async (c: Context) => {
        const id = c.req.param("id");

        const u = await repo.get(id);

        return c.json(u);
    };
};
