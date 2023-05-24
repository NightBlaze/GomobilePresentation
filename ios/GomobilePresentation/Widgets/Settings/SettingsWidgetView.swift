//
//  SettingsWidgetView.swift
//  GomobilePresentation
//
//  Created by Alexander Timonenkov on 24.05.2023.
//

import AutoLayoutSugar
import Engine
import FPSCounter
import UIKit

final class SettingsWidgetView: BaseWidgetView,
                                SettingswidgetDisplayProtocol,
                                FPSCounterDelegate {
    private lazy var widget = SettingswidgetCreate(self)
    private lazy var router = SettingsRouter(widget: self)

    private lazy var ruLocalizationButton: UIButton = {
        let view = UIButton(type: .system).prepareForAutoLayout()
        let action = UIAction { [weak self] _ in
            self?.widget?.changeToRuLocalizationAsync()
        }
        view.addAction(action, for: .touchUpInside)
        return view
    }()

    private lazy var enLocalizationButton: UIButton = {
        let view = UIButton(type: .system).prepareForAutoLayout()
        let action = UIAction { [weak self] _ in
            self?.widget?.changeToEnLocalizationAsync()
        }
        view.addAction(action, for: .touchUpInside)
        return view
    }()

    private let fpsLabel: UILabel = {
        let view = UILabel().prepareForAutoLayout()
        view.textColor = .green
        view.font = UIFont.systemFont(ofSize: 20)
        view.text = "00"
        return view
    }()

    private lazy var fpsCounter: FPSCounter = {
        let result = FPSCounter()
        result.delegate = self
        return result
    }()

    override init() {
        super.init()

        setupUI()
        setupLayout()
    }

    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    override func free() {
        widget?.free()
        widget = nil
    }

    // MARK: - Overrides

    override func viewControllerDidLoad() {
        super.viewControllerDidLoad()

        fpsCounter.startTracking()
        setupInitialData()
    }

    // MARK: - Private

    private func setupUI() {
        addSubview(ruLocalizationButton)
        addSubview(enLocalizationButton)
        addSubview(fpsLabel)
    }

    private func setupLayout() {
        ruLocalizationButton.leadingAnchor ~ leadingAnchor + 16
        ruLocalizationButton.topAnchor ~ topAnchor

        enLocalizationButton.leadingAnchor ~ leadingAnchor + 16
        enLocalizationButton.topAnchor ~ ruLocalizationButton.bottomAnchor + 16

        fpsLabel.trailingAnchor ~ trailingAnchor - 16
        fpsLabel.topAnchor ~ topAnchor + 16
    }

    private func setupInitialData() {
        guard let initialData = widget?.initialData() else { return }

        ruLocalizationButton.setTitle(initialData.ruLocalizationTitle, for: .normal)
        enLocalizationButton.setTitle(initialData.enLocalizationTitle, for: [])
    }

    // MARK: - Settingswidget

    func localizationDidChange(_ viewModel: SettingswidgetSettingLocalizationDidChangeViewModel?) {
        mainAsync { [weak self] in
            guard
                let self = self,
                let viewModel = viewModel
            else {
                return
            }

            ruLocalizationButton.setTitle(viewModel.ruLocalizationTitle, for: [])
            enLocalizationButton.setTitle(viewModel.enLocalizationTitle, for: [])
        }
    }

    // MARK: - FPSCounterDelegate

    func fpsCounter(_ counter: FPSCounter, didUpdateFramesPerSecond fps: Int) {
        fpsLabel.text = "\(fps)"
    }
}
