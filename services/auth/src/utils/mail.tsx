import { render } from "@react-email/render";
import nodemailer from "nodemailer";
import { EmailTemplate } from "@daveyplate/better-auth-ui/server";
import { type User } from "better-auth/types";

export const transporter = nodemailer.createTransport({
    service: "gmail",
    auth: {
        user: process.env.EMAIL_USERNAME as string,
        pass: process.env.EMAIL_PASSWORD as string,
    },
});

export const template = async (
    action: string,
    message: string,
    user: User,
    url: string,
    request?: Request,
): Promise<string> => {
    return await render(
        EmailTemplate({
            action,
            content: (
                <>
                    <p>{`Hello ${user.name},`}</p>

                    <p>{message}</p>
                </>
            ),
            heading: action,
            url,
        }),
    );
};
