import Link from 'next/link';
import MaxWidthWrapper from './MaxWidthWrapper';
import { buttonVariants } from './ui/button';
import { ArrowRight, Landmark } from 'lucide-react';

import MobileNav from './MobileNav';

export const Navbar = async () => {
  return (
    <nav className="sticky h-14 inset-x-0 top-0 z-30 w-full border-b border-gray-200 bg-white/75 backdrop-blur-lg transition-all">
      <MaxWidthWrapper>
        <div className="flex h-14 items-center justify-between border-b border-zinc-200">
          <Link href="/" className="flex z-40 font-semibold items-center">
            <Landmark className="h-4 w-4 mr-2" />
            <span>GoBank</span>
          </Link>

          {/* Mobile Navbar */}
          <MobileNav />

          {/* Desktop Navbar */}
          <div className="hidden items-center space-x-4 sm:flex">
            <>
              <Link
                href="/sign-in"
                className={buttonVariants({
                  variant: 'ghost',
                  size: 'sm',
                })}
              >
                Login
              </Link>
              <Link
                href="/sign-up"
                className={buttonVariants({
                  size: 'sm',
                })}
              >
                Sign Up <ArrowRight className="ml-1.5 h-5 w-5" />
              </Link>
            </>
          </div>
        </div>
      </MaxWidthWrapper>
    </nav>
  );
};
