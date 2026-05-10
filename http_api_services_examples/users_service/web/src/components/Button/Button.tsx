import { type ComponentProps } from 'react';
import { twMerge } from 'tailwind-merge';

type Variants = 'primary' | 'danger' | 'ghost-danger';

interface ButtonProps extends ComponentProps<'button'> {
  variant?: Variants;
}

export const Button = ({ variant = 'primary', ...props }: ButtonProps) => {
  return (
    <button
      disabled={props.disabled}
      {...props}
      className={twMerge(
        'disabled:text-gray-300 font-bold py-2 px-4 rounded cursor-pointer transition-colors',
        getVariantStyles(variant),
        props.className,
      )}
    >
      {props.children}
    </button>
  );
};

function getVariantStyles(variant: Variants) {
  switch (variant) {
    case 'primary':
      return 'bg-blue-500 hover:bg-blue-700 text-white disabled:bg-gray-400 disabled:text-gray-300 ';
    case 'danger':
      return 'bg-red-500 hover:bg-red-700 text-white disabled:bg-gray-400 disabled:text-gray-300';
    case 'ghost-danger':
      return 'bg-transparent hover:bg-red-500 hover:disabled:bg-transparent text-red-500 hover:text-white disabled:text-gray-300';
    default:
      throw new Error(`Invalid variant: ${variant satisfies never}`);
  }
}
