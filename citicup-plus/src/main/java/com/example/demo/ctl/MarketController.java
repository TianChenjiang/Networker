package com.example.demo.ctl;

import com.example.demo.model.Market;
import com.example.demo.repo.MarketRepo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/plus/markets")
public class MarketController {
    @Autowired
    private MarketRepo marketRepo;

    @GetMapping("/list")
    public List<Market> getAllMarket(@RequestParam(name = "pageSize") int size,
                                     @RequestParam(name = "pageNum") int pageNum) {
        PageRequest request = PageRequest.of(pageNum, size);
        return marketRepo.findAll(request).getContent();
    }
}
