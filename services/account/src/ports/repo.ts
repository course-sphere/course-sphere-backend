import { User } from "../domain/user";

export interface UserRepository {
    get(id: string): Promise<User>;
}
