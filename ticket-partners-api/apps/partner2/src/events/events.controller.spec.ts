import { Test, TestingModule } from '@nestjs/testing';
import { EventsController } from './events.controller';
import { EventsCoreModule } from '@app/core';

describe('EventsController', () => {
  let controller: EventsController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [EventsController],
      imports: [EventsCoreModule],
    }).compile();

    controller = module.get<EventsController>(EventsController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
