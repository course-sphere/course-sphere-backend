export type Role = "student" | "instructor" | "admin";

export type User = {
    email: string;
    name: string;
    username: string | null;
    displayUsername: string | null;
    emailVerified: boolean;
    image: string | null;
    role: Role;
    twoFactorEnabled: boolean | null;
    createdAt: Date;
    updatedAt: Date;
};
