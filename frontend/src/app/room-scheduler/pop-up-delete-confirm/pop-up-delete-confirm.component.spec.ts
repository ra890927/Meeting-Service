import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PopUpDeleteConfirmComponent } from './pop-up-delete-confirm.component';

describe('PopUpDeleteConfirmComponent', () => {
  let component: PopUpDeleteConfirmComponent;
  let fixture: ComponentFixture<PopUpDeleteConfirmComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PopUpDeleteConfirmComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(PopUpDeleteConfirmComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
