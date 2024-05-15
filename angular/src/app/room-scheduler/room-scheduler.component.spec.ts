import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RoomSchedulerComponent } from './room-scheduler.component';

describe('RoomSchedulerComponent', () => {
  let component: RoomSchedulerComponent;
  let fixture: ComponentFixture<RoomSchedulerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RoomSchedulerComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(RoomSchedulerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
