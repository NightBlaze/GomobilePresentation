//
//  ViewController.swift
//  GomobilePresentation
//
//  Created by Alexander Timonenkov on 23.05.2023.
//

import AutoLayoutSugar
import UIKit

class ViewController: BaseViewController {
    private lazy var settingsWidgetView: SettingsWidgetView = {
        let view = SettingsWidgetView().prepareForAutoLayout()
        view.setHostViewController(self)
        return view
    }()
    private lazy var feedWidgetView: FeedWidgetView = {
        let view = FeedWidgetView().prepareForAutoLayout()
        view.setHostViewController(self)
        return view
    }()

    override func viewDidLoad() {
        setupUI()
        setupLayout()

        super.viewDidLoad()
    }

    deinit {
        settingsWidgetView.free()
        feedWidgetView.free()
    }

    // MARK: - Private

    private func setupUI() {
        view.addSubview(settingsWidgetView)
        view.addSubview(feedWidgetView)
    }

    private func setupLayout() {
        settingsWidgetView.pin([.left, .right])
        settingsWidgetView.safeTopAnchor ~ view.safeTopAnchor + 100
        settingsWidgetView.heightAnchor ~ 80

        feedWidgetView.pin(excluding: .top)
        feedWidgetView.topAnchor ~ settingsWidgetView.bottomAnchor
    }
}
