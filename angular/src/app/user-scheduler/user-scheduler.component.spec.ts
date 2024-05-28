import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserSchedulerComponent } from './user-scheduler.component';

describe('UserSchedulerComponent', () => {
  let component: UserSchedulerComponent;
  let fixture: ComponentFixture<UserSchedulerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [UserSchedulerComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(UserSchedulerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
