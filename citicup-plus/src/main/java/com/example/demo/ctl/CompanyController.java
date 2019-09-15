package com.example.demo.ctl;

import com.example.demo.model.Company;
import com.example.demo.model.Market;
import com.example.demo.repo.CompanyRepo;
import com.example.demo.repo.MarketRepo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.stream.Collectors;

@RestController
@RequestMapping("/api/plus/companies")
public class CompanyController {
    @Autowired
    private CompanyRepo companyRepo;
    @Autowired
    private MarketRepo marketRepo;

    @GetMapping
    public List<Company> getAllCompanies(@RequestParam(name = "pageSize") int size,
                                         @RequestParam(name = "pageNum") int pageNum) {
        return companyRepo.findAll(PageRequest.of(pageNum, size))
                .getContent()
                .stream()
                .peek(company ->
                        {
                            List<Market> markets = marketRepo.findMarketsByTsCode(company.getTsCode());
                            if (markets.size() != 0) {
                                company.setMarket(markets.get(0));
                            }
                        }
                )
                .collect(Collectors.toList());
    }

}
