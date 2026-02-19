import { DBUserRepository } from "./adapters/repo/user.db";
import { serve } from "./transport/http/server";

const app = serve(new DBUserRepository());

export default app;
