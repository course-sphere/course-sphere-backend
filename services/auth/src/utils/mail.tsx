import { render } from "@react-email/render";
import { type User } from "better-auth/types";
import nodemailer from "nodemailer";
import { EmailTemplate } from "@daveyplate/better-auth-ui/server";

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
    const builder = new URL(url);
    builder.host = new URL(request?.headers.get("origin") as string).host;

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
