import { type AxiosResponse, type InternalAxiosRequestConfig } from 'axios'
import { afterEach, describe, expect, it } from 'vitest'

import { http } from '@/shared/lib/http'

function resolveResponse(config: InternalAxiosRequestConfig): Promise<AxiosResponse> {
  return Promise.resolve({
    config,
    data: null,
    headers: {},
    status: 200,
    statusText: 'OK',
  })
}

describe('http', () => {
  afterEach(() => {
    // biome-ignore lint/suspicious/noDocumentCookie: the interceptor reads the browser cookie API.
    document.cookie = 'csrf_token=; Max-Age=0; path=/'
  })

  it('forwards the csrf token cookie as a request header', async () => {
    // biome-ignore lint/suspicious/noDocumentCookie: the interceptor reads the browser cookie API.
    document.cookie = 'csrf_token=token-value; path=/'

    const response = await http.get('/csrf-test', { adapter: resolveResponse })

    expect(response.config.headers['X-CSRF-Token']).toBe('token-value')
  })
})
