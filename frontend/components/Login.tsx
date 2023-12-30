'use client';

import * as z from 'zod';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from './ui/form';
import { Input } from './ui/input';
import { Button } from './ui/button';
import Link from 'next/link';

const formSchema = z.object({
  emailAddress: z.string().email(),
  password: z.string().min(3),
});

const Login = () => {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      emailAddress: '',
      password: '',
    },
  });

  const onSubmit = (values: z.infer<typeof formSchema>) => {
    console.log({ values });
  };

  return (
    <div className="relative flex flex-col justify-center items-center min-h-screen overflow-hidden">
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="max-w-md w-full flex flex-col gap-4 border p-10 rounded-md bg-slate-50"
        >
          <h2 className="mb-2 text-2xl">Login</h2>
          <FormField
            control={form.control}
            name="emailAddress"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input type="email" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input type="password" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <div className="flex justify-end">
            <Link href="/" className="text-xs hover:text-slate-600">
              Forgot Password
            </Link>
          </div>
          <span>
            Don&apos;t have an account?{' '}
            <Link
              href="/sign-up"
              className="text-gray-700 underline underline-offset-2"
            >
              sign up
            </Link>
          </span>
          <Button type="submit">Continue</Button>
        </form>
      </Form>
    </div>
  );
};

export default Login;
