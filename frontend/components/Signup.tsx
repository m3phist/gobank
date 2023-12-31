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
import { useToast } from './ui/use-toast';
import Link from 'next/link';
import axios from 'axios';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { authUrl } from '@/utils/network';

const formSchema = z
  .object({
    emailAddress: z.string().email(),
    password: z.string().min(3),
    passwordConfirm: z.string(),
  })
  .refine(
    (data) => {
      return data.password === data.passwordConfirm;
    },
    {
      message: 'Password do not match',
      path: ['passwordConfirm'],
    }
  );

const Signup = () => {
  const { toast } = useToast();
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      emailAddress: '',
      password: '',
      passwordConfirm: '',
    },
  });

  const onSubmit = (values: z.infer<typeof formSchema>) => {
    const { emailAddress, password } = values;
    let arg = {
      email: emailAddress,
      password: password,
    };

    console.log(arg);
    setIsLoading(true);
    console.log(authUrl.register);

    axios
      .post(authUrl.register, arg)
      .then(() => {
        toast({
          variant: 'success',
          description: 'Your account has been created.',
        });

        router.push('/sign-in');
      })
      .catch((error) => {
        toast({
          variant: 'destructive',
          description: 'Something went wrong.',
        });
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  return (
    <div className="relative flex flex-col justify-center items-center min-h-screen overflow-hidden">
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="max-w-md w-full flex flex-col gap-4 border p-10 rounded-md bg-slate-50"
        >
          <h2 className="mb-2 text-2xl">Create an account</h2>
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
          <FormField
            control={form.control}
            name="passwordConfirm"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Confirm Password</FormLabel>
                <FormControl>
                  <Input type="password" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <span>
            Already have an account?{' '}
            <Link
              href="/sign-in"
              className="text-gray-700 underline underline-offset-2"
            >
              sign in
            </Link>
          </span>
          <Button type="submit" disabled={isLoading}>
            Register
          </Button>
        </form>
      </Form>
    </div>
  );
};

export default Signup;
