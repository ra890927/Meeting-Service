import { TestBed } from '@angular/core/testing';

import { RoomAPIService } from './room-api.service';

describe('RoomAPIService', () => {
  let service: RoomAPIService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RoomAPIService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
