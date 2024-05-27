import { TestBed } from '@angular/core/testing';

import { AuthHttpInterceptorServiceService } from './auth-http-interceptor-service.service';

describe('AuthHttpInterceptorServiceService', () => {
  let service: AuthHttpInterceptorServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AuthHttpInterceptorServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
