/*
Copyright (C) 2023-2026 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/
import { CopyButton } from '@/components/copy-button'
import { Input } from '@/components/ui/input'

const API_BASE_URL = 'https://mokunx.com/v1'
const API_ENDPOINT = '/responses/compact'

export function ApiBaseUrlHint() {
  return (
    <div className='flex w-full max-w-xl flex-col items-start gap-3'>
      <p className='text-muted-foreground/80 text-sm md:text-base'>
        多模型统一接入，只需将基址替换为：
      </p>
      <div className='flex w-full items-center gap-2 rounded-full border border-border/40 bg-muted/20 p-1.5 backdrop-blur-xs'>
        <Input
          aria-label='Base URL'
          readOnly
          value={API_BASE_URL}
          className='h-9 rounded-full border-0 bg-transparent px-3 shadow-none focus-visible:ring-0'
        />
        <span className='text-muted-foreground/80 shrink-0 text-sm'>
          {API_ENDPOINT}
        </span>
        <CopyButton
          value={API_BASE_URL}
          variant='ghost'
          size='icon'
          tooltip='复制 Base URL'
          successTooltip='已复制'
          className='text-muted-foreground hover:bg-muted hover:text-foreground size-9 rounded-full'
        />
      </div>
    </div>
  )
}
