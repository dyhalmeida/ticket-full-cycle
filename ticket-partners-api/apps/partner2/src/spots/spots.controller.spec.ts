import { Test, TestingModule } from '@nestjs/testing';
import { SpotsController } from './spots.controller';
import { SpotsCoreModule } from '@app/core';

describe('SpotsController', () => {
  let controller: SpotsController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      imports: [SpotsCoreModule],
      controllers: [SpotsController],
    }).compile();

    controller = module.get<SpotsController>(SpotsController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
